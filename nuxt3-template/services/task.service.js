export default function useTaskService() {
    const { $axios } = useNuxtApp();
  
    const getAllTasks = async () => {
      try {
        const response = await $axios.get('/tasks');
        return response.data;
      } catch (error) {
        console.error('Error al obtener tareas:', error);
        throw error;
      }
    };
  
    const createTask = async (taskData) => {
      try {
        const response = await $axios.post('/tasks', taskData);
        return response.data;
      } catch (error) {
        console.error('Error al crear tarea:', error);
        throw error;
      }
    };
  
    const markAsCompleted = async (taskId) => {
      try {
        const response = await $axios.put(`/tasks/${taskId}/complete`);
        return response.data;
      } catch (error) {
        console.error('Error al marcar tarea como completada:', error);
        throw error;
      }
    };
  
    return {
      getAllTasks,
      createTask,
      markAsCompleted
    };
  }
  