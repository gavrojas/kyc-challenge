<script setup lang="ts">
import { onMounted, computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useValidations } from '@/services/validation'
import { useDocumentStore } from '@/stores/documents'
import { useNavigationStore } from '@/stores/navigation'
import type { ValidationResultInfo } from '@/types'

const { getValidationsResult } = useValidations()
const router = useRouter()
const documentStore = useDocumentStore()
const navigationStore = useNavigationStore()
const loading = ref<boolean>(false)

const validationResponse = computed<ValidationResultInfo | null>(
  () => documentStore.ValidationResultInfo
)

const processingTime = computed(() => {
  const startDateStr = validationResponse.value?.processing_start_date
  const finishDateStr = validationResponse.value?.processing_finish_date

  if (startDateStr && finishDateStr) {
    const startDate = new Date(startDateStr)
    const finishDate = new Date(finishDateStr)

    if (!isNaN(startDate.getTime()) && !isNaN(finishDate.getTime())) {
      const differenceInMilliseconds = finishDate.getTime() - startDate.getTime()
      return Math.floor(differenceInMilliseconds / 1000)
    }
  }

  return 0
})

const handleBack = () => {
  navigationStore.previousStep()
  router.push('/onboarding-step-3')
}

const checkValidationStatus = async () => {
  while (validationResponse.value?.validation_status === 'pending') {
    loading.value = true
    await new Promise((resolve) => setTimeout(resolve, 25000)) // 25 seconds
    documentStore.ValidationResultInfo = await getValidationsResult()
  }
  loading.value = false
}

onMounted(async () => {
  loading.value = true
  documentStore.ValidationResultInfo = await getValidationsResult()
  await checkValidationStatus()
})
</script>

<template>
  <v-container class="p-4">
    <v-card>
      <v-card-title class="display-1"> Validation Results </v-card-title>
      <v-card-text>
        <div v-if="loading" class="d-flex justify-center align-center" style="height: 100px">
          <v-progress-circular indeterminate color="primary"></v-progress-circular>
          <span class="ml-2"
            >Loading validation information<br />
            Please wait...</span
          >
        </div>
        <p v-else>Below you will find the results of your document validation.</p>
      </v-card-text>
      <v-card-text v-if="validationResponse">
        <v-card-title class="display-2"> Validation Details </v-card-title>
        <p><strong>Status:</strong> {{ validationResponse.validation_status }}</p>
        <p v-if="validationResponse.failure_status">
          <strong>Failure status:</strong> {{ validationResponse.failure_status }}
        </p>
        <p v-if="validationResponse.declined_reason">
          <strong>Reason failure:</strong> {{ validationResponse.declined_reason }}
        </p>
        <p><strong>Processing Time (seconds):</strong> {{ processingTime }}</p>
        <v-row>
          <v-col cols="12" md="6">
            <v-card-title class="display-2"> User Data </v-card-title>
            <p><strong>Gender:</strong> {{ validationResponse.details.document_details.gender }}</p>
            <p><strong>Height:</strong> {{ validationResponse.details.document_details.height }}</p>
            <p><strong>RH:</strong> {{ validationResponse.details.document_details.rh }}</p>
            <p>
              <strong>Country:</strong> {{ validationResponse.details.document_details.country }}
            </p>
            <p>
              <strong>Birth Place:</strong>
              {{ validationResponse.details.document_details.birth_place }}
            </p>
            <p>
              <strong>Birth Date:</strong>
              {{ validationResponse.details.document_details.date_of_birth }}
            </p>
            <p><strong>Doc ID:</strong> {{ validationResponse.details.document_details.doc_id }}</p>
            <p>
              <strong>Expedition Place:</strong>
              {{ validationResponse.details.document_details.expedition_place }}
            </p>
            <p>
              <strong>Expedition Date:</strong>
              {{ validationResponse.details.document_details.issue_date }}
            </p>
            <p>
              <strong>Document Version:</strong>
              {{ validationResponse.details.document_details.document_version }}
            </p>
            <p>
              <strong>National Registrar:</strong>
              {{ validationResponse.details.document_details.national_registrar }}
            </p>
            <p>
              <strong>Production Data:</strong>
              {{ validationResponse.details.document_details.production_data }}
            </p>
          </v-col>

          <v-col cols="12" md="6">
            <v-card-title class="display-2"> Images uploaded </v-card-title>
            <div class="d-flex">
              <v-img
                v-if="validationResponse.user_response.input_files[0]"
                :src="validationResponse.user_response.input_files[0]"
                class="mr-4"
                width="150"
              />
              <v-img
                v-if="validationResponse.user_response.input_files[1]"
                :src="validationResponse.user_response.input_files[1]"
                width="150"
              />
            </div>
          </v-col>
        </v-row>
      </v-card-text>

      <v-card-text v-else>
        <p>No validation data available.</p>
      </v-card-text>
      <v-card-actions>
        <v-btn @click="handleBack" color="secondary">Back</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>
