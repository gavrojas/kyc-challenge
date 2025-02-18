<script setup lang="ts">
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { useDocumentStore } from '@/stores/documents'
import { useNavigationStore } from '@/stores/navigation'
import { useValidations } from '@/services/validation'

const router = useRouter()
const documentStore = useDocumentStore()
const navigationStore = useNavigationStore()
const alertVisible = ref(false)
const { uploadImages } = useValidations()

const handleBack = () => {
  navigationStore.previousStep()
  alertVisible.value = false
  router.push('/onboarding-step-2')
}

const handleFrontUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    documentStore.frontImage = target.files[0]
  }
}

const handleBackUpload = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    documentStore.backImage = target.files[0]
  }
}

const putImages = async () => {
  if (documentStore.frontImage) {
    documentStore.loading = true
    try {
      const response = await uploadImages('front_url', documentStore.frontImage)
      documentStore.validationResult = response
    } catch (error) {
      console.error('Failed to upload front image:', error)
    } finally {
      documentStore.loading = false
    }
  }

  if (documentStore.backImage) {
    documentStore.loading = true
    try {
      const response = await uploadImages('reverse_url', documentStore.backImage)
      documentStore.validationResult = response
    } catch (error) {
      console.error('Failed to upload back image:', error)
    } finally {
      documentStore.loading = false
    }
  }

  if (documentStore.validationResult) {
    alertVisible.value = true
    // wait 3 segundos
    setTimeout(() => {
      router.push('/results')
    }, 3000)
  }
}
</script>

<template>
  <v-container class="p-4">
    <!-- Barra de progreso -->
    <v-progress-linear
      :value="navigationStore.progress"
      height="5"
      color="primary"
      v-if="navigationStore.progress"
    />
    <v-card class="p-4">
      <v-card-title> Step 3 of 3: Upload your images </v-card-title>
      <v-card-text>
        <v-file-input
          label="Upload front image"
          @change="handleFrontUpload"
          variant="underlined"
          accept="image/*"
        ></v-file-input>
        <v-file-input
          label="Upload back image"
          @change="handleBackUpload"
          variant="underlined"
          accept="image/*"
        ></v-file-input>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="handleBack" color="secondary">Back</v-btn>
        <v-btn :loading="documentStore.loading" @click="putImages" color="primary">Validate</v-btn>
      </v-card-actions>
      <v-alert v-if="documentStore.validationResult && alertVisible" type="success">
        Result: {{ documentStore.validationResult.message }}
      </v-alert>
    </v-card>
  </v-container>
</template>
