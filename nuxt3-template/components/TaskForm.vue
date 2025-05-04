<template>
    <div class="modal-backdrop">
      <div class="modal-container">
        <h2 class="text-xl font-bold mb-4">Crear Nueva Tarea</h2>
  
        <form @submit.prevent="submitTask">
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700">Título</label>
            <input
              v-model="title"
              type="text"
              class="mt-1 block w-full border border-gray-300 rounded px-3 py-2"
              placeholder="Ej: Comprar leche"
            />
            <p v-if="errors.title" class="text-red-600 text-sm mt-1">{{ errors.title }}</p>
          </div>
  
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700">Descripción</label>
            <textarea
              v-model="description"
              class="mt-1 block w-full border border-gray-300 rounded px-3 py-2"
              placeholder="Ej: Ir al supermercado y comprar leche descremada"
            ></textarea>
            <p v-if="errors.description" class="text-red-600 text-sm mt-1">{{ errors.description }}</p>
          </div>
  
          <div class="flex justify-end space-x-2">
            <button
              type="button"
              @click="emit('cancel')"
              class="btn-cancel"
            >
              Cancelar
            </button>
  
            <button
              type="submit"
              class="btn-create"
            >
              Crear tarea
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, reactive } from 'vue'
  import useTaskService from '@/services/task.service'
  
  const { createTask } = useTaskService()
  
  const emit = defineEmits(['created', 'cancel'])
  
  const title = ref('')
  const description = ref('')
  const errors = reactive({ title: '', description: '' })
  
  async function submitTask() {
    errors.title = ''
    errors.description = ''
  
    if (!title.value.trim()) {
      errors.title = 'El título es obligatorio.'
      return
    }
  
    if (description.value.length < 5) {
      errors.description = 'La descripción debe tener al menos 5 caracteres.'
      return
    }
  
    try {
      await createTask({
      title: title.value,
      description: description.value
    })
  
      title.value = ''
      description.value = ''
      emit('created')
      alert('Tarea creada correctamente')
    } catch (error) {
      console.error('Error al crear la tarea:', error)
      alert('No se pudo crear la tarea.')
    }
  }
  </script>
  
  <style scoped>
    input,
    textarea {
    width: 100%;
    padding: 10px 14px;
    font-size: 14px;
    border: 1px solid #d1d5db; 
    border-radius: 6px;
    outline: none;
    transition: border-color 0.2s, box-shadow 0.2s;
    background-color: #f9fafb;
    resize: none; 
    box-sizing: border-box;
    }

    textarea {
    height: 100px; 
    overflow-y: auto; 
    }

    input:focus,
    textarea:focus {
    border-color: #3b82f6; 
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.2);
    background-color: white;
    }

    label {
        margin-bottom: 6px;
        display: inline-block;
        color: #374151;
        font-size: 16px;
        font-weight: 600;
    }


  
    .modal-backdrop {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.3);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 50;
    }
  
    .modal-container {
        background: #fff;
        border-radius: 16px;
        padding: 32px;
        width: 90%;
        max-width: 600px;
        max-height: 90vh;
        box-shadow: 0 20px 50px rgba(0, 0, 0, 0.2);
        overflow-y: auto;
        z-index: 100;
    }

    .modal-title {
    text-align: center;
    font-size: 22px;
    font-weight: 700;
    margin-bottom: 20px;
    color: #111827;
    position: relative;
    }

    .close-button {
    position: absolute;
    top: 14px;
    right: 14px;
    background: none;
    border: none;
    font-size: 22px;
    font-weight: bold;
    color: #6b7280;
    cursor: pointer;
    transition: color 0.2s ease;
    }
    .close-button:hover {
    color: #111827;
    }

    /* Botones */
    button[type="submit"] {
    background-color: #3b82f6;
    color: white;
    font-weight: 600;
    padding: 8px 16px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s ease;
    }
    button[type="submit"]:hover {
    background-color: #2563eb;
    }

    button[type="button"] {
    background-color: #e5e7eb;
    color: #374151;
    font-weight: 500;
    padding: 8px 16px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s ease;
    }
    button[type="button"]:hover {
    background-color: #d1d5db;
    }

    .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 12px;
    margin-top: 20px;
    }

    .btn-cancel {
    background-color: #e5e7eb;
    color: #374151;
    font-weight: 500;
    padding: 8px 16px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s ease;
    }
    .btn-cancel:hover {
    background-color: #d1d5db;
    }

    .btn-create {
    background-color: #3b82f6;
    color: white;
    font-weight: 600;
    padding: 8px 16px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s ease;
    }
    .btn-create:hover {
    background-color: #2563eb;
    }

    button[type="submit"],
    button[type="button"],
    .btn-create,
    .btn-cancel {
        font-size: 15px;
        padding: 10px 20px;
    }



  </style>
  