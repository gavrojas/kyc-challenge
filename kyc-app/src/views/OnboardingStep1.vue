<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useNavigationStore } from '@/stores/navigation'
import { useDocumentStore } from '@/stores/documents'
import { useValidations } from '@/services/validation'
import type { Country, DocType } from '@/types'

const { createValidation, getInfoCountries } = useValidations()
const accepted = ref(false)

const navigationStore = useNavigationStore()
const documentStore = useDocumentStore()

const countries = Object.keys(documentStore.infoCountries) as Country[]
const country = ref<Country | null>(null)
const docTypes = ref<DocType[]>([])
const docType = ref<DocType | null>(null)

const router = useRouter()

const handleNext = async () => {
  documentStore.setUserAuthorization(accepted.value)
  if (accepted.value && country.value) {
    navigationStore.nextStep()
    router.push('/onboarding-step-2')
  } else {
    alert('Please accept and select a country')
  }
  await createValidation()
}

const handleCountryChange = (selectedCountry: Country) => {
  documentStore.setCountry(selectedCountry) // Guardar el país en el store
  docTypes.value = documentStore.infoCountries[selectedCountry]?.documentTypes || []
}

const handleDocType = (selectedDocType: DocType) => {
  documentStore.setDocType(selectedDocType)
}

watch(country, (newCountry) => {
  if (newCountry) {
    handleCountryChange(newCountry)
  }
})

onMounted(async () => {
  await getInfoCountries()
})
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
      <v-card-title> Step 1 of 3: Authorization </v-card-title>
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
        <v-select
          v-model="docType"
          :items="docTypes"
          label="Select your document type"
          variant="underlined"
          item-text="name"
          item-value="value"
          @change="handleDocType($event)"
        />
      </v-card-text>
      <v-card-actions>
        <v-btn @click="handleNext" color="primary">Continue</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
