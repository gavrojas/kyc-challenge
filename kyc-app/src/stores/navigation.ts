import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useNavigationStore = defineStore('navigation', () => {
  const step = ref(1) // Paso actual (1, 2 o 3)
  const progress = ref(33) // Valor del progreso (33, 66, 100)

  const nextStep = () => {
    if (step.value < 3) {
      step.value++
      updateProgress()
    }
  }

  const previousStep = () => {
    if (step.value > 1) {
      step.value--
      updateProgress()
    }
  }

  const updateProgress = () => {
    if (step.value === 1) {
      progress.value = 33
    } else if (step.value === 2) {
      progress.value = 66
    } else if (step.value === 3) {
      progress.value = 100
    }
  }

  return { step, progress, nextStep, previousStep }
})
