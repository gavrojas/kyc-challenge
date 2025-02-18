import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { SessionData, Country, ValidationResultInfo, DocType, InfoCountries } from '@/types'

export const useDocumentStore = defineStore('documents', () => {
  const frontImage = ref<File | null>(null)
  const backImage = ref<File | null>(null)
  const validationResult = ref<{ message: string } | null>(null)
  const loading = ref<boolean>(false)
  const sessionData = ref<SessionData[]>([])
  const ValidationResultInfo = ref<ValidationResultInfo | null>(null)

  const infoCountries: InfoCountries = {
    Colombia: { documentTypes: [] },
    Chile: { documentTypes: [] },
    Mexico: { documentTypes: [] },
    Peru: { documentTypes: [] },
    Brasil: { documentTypes: [] },
    'Costa Rica': { documentTypes: [] },
    Other: { documentTypes: [] }
  }

  const type = ref<string>('document-validation')
  const timeout = ref(900)
  const documentType = ref<DocType>('national-id')
  const selectedCountry = ref<Country>('Colombia')
  const userAuthorization = ref<boolean>(false)

  const setCountry = (country: Country) => {
    selectedCountry.value = country
  }

  const setDocType = (docType: DocType) => {
    documentType.value = docType
  }

  const setUserAuthorization = (value: boolean) => {
    userAuthorization.value = value
  }

  const updateInfoCountries = (newInfoCountries: InfoCountries) => {
    Object.assign(infoCountries, newInfoCountries)
  }

  return {
    frontImage,
    backImage,
    validationResult,
    type,
    timeout,
    documentType,
    sessionData,
    loading,
    userAuthorization,
    selectedCountry,
    ValidationResultInfo,
    infoCountries,
    setCountry,
    setDocType,
    setUserAuthorization,
    updateInfoCountries
  }
})
