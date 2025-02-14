<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useDocumentStore } from '@/stores/documents';
import { useNavigationStore } from '@/stores/navigation';

const router = useRouter();
const documentsStore = useDocumentStore();
const navigationStore = useNavigationStore();

const handleBack = () => {
  navigationStore.previousStep();
  router.push('/onboarding-step-2'); 
};

const handleFrontUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files) {
    documentsStore.frontImage = target.files[0];
  }
};

const handleBackUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files) {
    documentsStore.backImage = target.files[0];
  }
};
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
      <v-card-title>
        Step 3 of 3: Upload your images
      </v-card-title>
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
        <v-btn :loading="documentsStore.loading" @click="documentsStore.uploadImages" color="primary">Validate</v-btn>
      </v-card-actions>
      <v-alert v-if="documentsStore.validationResult" type="success">
        Result: {{ documentsStore.validationResult.message }}
      </v-alert>
    </v-card>
  </v-container>
</template>
