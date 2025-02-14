<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useNavigationStore } from '@/stores/navigation';

const router = useRouter();
const store = useNavigationStore();

const handleBack = () => {
  store.previousStep();
  router.push('/onboarding-step-1'); // Volver al paso 1
};

const handleNext = () => {
  store.nextStep();
  router.push('/onboarding-step-3');
  console.log(store.progress);
};
</script>

<template>
  <v-container class="p-4">
    <!-- Barra de progreso -->
    <v-progress-linear 
      :value="store.progress"
      height="5"
      color="primary"
      v-if="store.progress"
    />
    <v-card class="p-4">
      <v-card-title>
        Step 2 of 3: How it works
      </v-card-title>
      <v-card-text>
        <p>
          You will be asked to upload two images of your document: the front and back.
          Please ensure you have access to your camera and multimedia.
        </p>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="handleBack" color="secondary">Back</v-btn>
        <v-btn @click="handleNext" color="primary">Continue</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
