<template>
    <div class="task-wrapper">
      <div class="header">
        <h2>Lista de Tareas</h2>
        <div class="header-actions">
            <button class="btn-new" @click="showForm = true">+ Nueva tarea</button>
            <TaskForm v-if="showForm" @created="onTaskCreated" @cancel="showForm = false" />

          <div class="filters">
            <button :class="{ active: filter === 'all' }" @click="filter = 'all'">Todas</button>
            <button :class="{ active: filter === 'completed' }" @click="filter = 'completed'">Completadas</button>
            <button :class="{ active: filter === 'pending' }" @click="filter = 'pending'">Pendientes</button>
          </div>
        </div>
      </div>
  
      <div v-if="filteredTasks.length > 0" class="table-container">
        <table>
          <thead>
            <tr>
              <th>Título</th>
              <th>Descripción</th>
              <th>Estado</th>
              <th>Creación</th>
              <th>Acción</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="task in filteredTasks" :key="task.id" :class="{ completed: task.completed }">
              <td>{{ task.title }}</td>
              <td>{{ task.description }}</td>
              <td>
                <span :class="task.completed ? 'estado completado' : 'estado pendiente'">
                  {{ task.completed ? 'Completada' : 'Pendiente' }}
                </span>
              </td>
              <td class="fecha">{{ formatDate(task.createdAt) }}</td>
              <td class="acciones">
                <button
                  v-if="!task.completed"
                  @click="completeTask(task.id)"
                >
                  Marcar completada
                </button>
                <span v-else class="check">✔</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
  
      <p v-else class="mensaje-vacio">No hay tareas para mostrar.</p>
    </div>
  </template>
  
  <script setup>
  import TaskForm from './TaskForm.vue'
  import { ref, computed, defineProps, defineEmits } from 'vue'
  import useTaskService from '@/services/task.service'

  const { markAsCompleted } = useTaskService()
  
  const showForm = ref(false)

  const props = defineProps({
    tasks: {
      type: Array,
      required: true
    }
  })
  const emit = defineEmits(['refresh'])
  
  const filter = ref('all')
  
  const filteredTasks = computed(() => {
    if (filter.value === 'completed') return props.tasks.filter(t => t.completed)
    if (filter.value === 'pending') return props.tasks.filter(t => !t.completed)
    return props.tasks
  })
  
    async function completeTask(id) {
    try {
      await markAsCompleted(id)
      emit('refresh')
    } catch (err) {
      alert('No se pudo marcar como completada.')
      console.error(err)
    }
  }


  function onTaskCreated() {
  showForm.value = false
  emit('refresh')
  }

  
  function formatDate(dateStr) {
    const date = new Date(dateStr)
    return date.toLocaleString('es-CL')
  }
  </script>
  
  <style scoped>
    .task-wrapper {
    background: #fff;
    border-radius: 16px;
    padding: 32px;
    box-shadow: 0 6px 18px rgba(0, 0, 0, 0.06);
    max-width: 90%;
    height: auto;
    margin: 40px auto;
    font-size: 16px;
    }

    thead th {
    font-size: 16px;
    }

    tbody td {
    font-size: 15px;
    }


  
    .header {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-bottom: 24px;
    border-bottom: 1px solid #e2e8f0;
    padding-bottom: 16px;
    }
  
    .header h2 {
    font-size: 32px;
    font-weight: 700;
    color: #1f2937;
    margin: 0;
    }
  
    .header-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
    gap: 12px;
    }
  
    .btn-new {
    background-color: #10b981;
    color: white;
    padding: 10px 20px;
    font-size: 16px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s ease;
    }

    .btn-new:hover {
    background-color: #059669;
    }
  
    .filters {
    display: flex;
    gap: 8px;
    }
  
    .filters button {
    padding: 6px 14px;
    border-radius: 6px;
    border: 1px solid #ccc;
    background-color: #f3f4f6;
    cursor: pointer;
    transition: all 0.2s ease;
    font-size: 14px;
    }
  
    .filters button.active,
    .filters button:hover {
    background-color: #2563eb;
    color: #fff;
    border-color: #2563eb;
    }
  
    .table-container {
    overflow-x: auto;
    }
  
    table {
    width: 100%;
    border-collapse: collapse;
    font-size: 14px;
    }
  
    thead th {
    background-color: #f9f9f9;
    text-align: left;
    padding: 10px;
    border-bottom: 2px solid #ddd;
    }
  
    tbody td {
    padding: 10px;
    border-bottom: 1px solid #eee;
    }
  
    tbody tr:hover {
    background-color: #f6f6f6;
    }
  
    tbody tr.completed {
    background-color: #e6ffec;
    }
  
    .estado.completado {
    color: #28a745;
    font-weight: 600;
    }
  
    .estado.pendiente {
    color: #ffc107;
    font-weight: 600;
    }
  
    .fecha {
    color: #666;
    }
  
    .acciones button {
    padding: 5px 10px;
    background-color: #007bff;
    border: none;
    color: white;
    font-size: 13px;
    border-radius: 4px;
    cursor: pointer;
    }

    .acciones button:hover {
    background-color: #0056b3;
    }
  
    .acciones .check {
    font-size: 20px;
    color: #28a745;
    }
  
    .mensaje-vacio {
    text-align: center;
    color: #777;
    margin-top: 20px;
    }
  </style>
  