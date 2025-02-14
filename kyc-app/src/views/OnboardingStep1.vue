<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useNavigationStore } from '@/stores/navigation';

const accepted = ref(false);
const country = ref('');
const countries = ['Colombia', 'Chile', 'México', 'Peru', 'Brasil', 'Costa Rica'];

const store = useNavigationStore();
const router = useRouter();

const handleNext = () => {
  if (accepted.value && country.value) {
    store.nextStep();
    router.push('/onboarding-step-2'); 
    console.log(store.progress);
  } else {
    alert('Please accept and select a country');
  }
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
        Step 1 of 3: Authorization
      </v-card-title>
      <v-card-text>
        <v-checkbox v-model="accepted" label="I accept the document validation" />
        <v-select
          v-model="country"
          :items="countries"
          label="Select your country"
          variant="underlined"
          item-text="name"
          item-value="value"
        />
      </v-card-text>
      <v-card-actions>
        <v-btn @click="handleNext" color="primary">Continue</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
