<template>
  <div>
    <div v-if="isLoading" class="text-gray-500 p-6">Cargando tareas...</div>
    <TaskList v-else :tasks="tasks" @refresh="handleRefresh" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import TaskList from '~/components/TaskList.vue'
import useTaskService from '@/services/task.service'

const { getAllTasks } = useTaskService()

const tasks = ref([])
const isLoading = ref(true)

const fetchTasks = async () => {
  isLoading.value = true
  try {
    const response = await getAllTasks()
    tasks.value = response
  } catch (error) {
    console.error('Error al obtener tareas:', error)
  } finally {
    isLoading.value = false
  }
}


onMounted(fetchTasks)

const handleRefresh = () => fetchTasks()
</script>
