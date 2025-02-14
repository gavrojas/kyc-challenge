import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useDocumentStore = defineStore('documents', () => {
  const frontImage = ref<File | null>(null);
  const backImage = ref<File | null>(null);
  const validationResult = ref<{ message: string } | null>(null);
  const loading = ref<boolean>(false);

  const uploadImages = async (): Promise<void> => {
    if (!frontImage.value || !backImage.value) {
      alert('Please upload both images');
      return;
    }

    loading.value = true;
    const formData = new FormData();
    formData.append('front', frontImage.value);
    formData.append('back', backImage.value);

    try {
      /*
      Insertar integración API upload
      */
      // const result = await response.json();
      // validationResult.value = result;
    } catch (error) {
      console.error('Validation failed:', error);
    } finally {
      loading.value = false;
    }
  };

  return { frontImage, backImage, validationResult, loading, uploadImages };
});
