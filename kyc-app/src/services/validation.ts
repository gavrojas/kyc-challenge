import { apiValidate } from '@/services/utils'
import { useDocumentStore } from '@/stores/documents'
import type { Country } from '@/types'

export function useValidations() {
  const documentStore = useDocumentStore()

  async function getInfoCountries() {
    try {
      const response = await apiValidate(`/get-config`, {
        method: 'GET',
        upload: false
      })

      documentStore.updateInfoCountries(response)
      // console.log("session data", documentStore.infoCountries);
    } catch (error) {
      console.error('Get countries info failed:', error)
    } finally {
      documentStore.loading = false
    }
  }

  async function createValidation() {
    const countryCodeMap: Record<Country, string> = {
      Colombia: 'CO',
      'Costa Rica': 'CR',
      Brasil: 'BR',
      Chile: 'CH',
      Peru: 'PE',
      Mexico: 'MX',
      Other: 'ALL'
    }

    const countryCode = countryCodeMap[documentStore.selectedCountry]

    const body = {
      type: documentStore.type,
      user_authorized: documentStore.userAuthorization,
      country: countryCode,
      timeout: documentStore.timeout,
      document_type: documentStore.documentType
    }

    documentStore.loading = true

    try {
      const response = await apiValidate(`/create`, {
        method: 'POST',
        data: body,
        upload: false
      })

      documentStore.sessionData = [response]
      console.log('session data', documentStore.sessionData)

      document.cookie = `account_id=${documentStore.sessionData[0].account_id}; path=/;`
    } catch (error) {
      console.error('Validation failed:', error)
    } finally {
      documentStore.loading = false
    }
  }

  async function uploadImages(imageType: 'front_url' | 'reverse_url', file: File) {
    if (documentStore.sessionData.length <= 0) {
      // ** lógica para recrear sesión
      throw new Error('No session data available')
    }

    const accountId = getCookie('account_id') || ''
    if (!accountId) {
      console.log('No account ID cookie found.')
    }

    const formData = new FormData()
    formData.append('file', file)

    const params: Record<string, string> = {
      image_type: imageType,
      account_id: accountId
    }

    try {
      const query = new URLSearchParams(params).toString()
      const response = await apiValidate(`/put-file?${query}`, {
        method: 'POST',
        data: formData,
        upload: true
      })
      // console.log(response);

      return response
    } catch (error) {
      console.error('Images upload failed:', error)
      throw error
    }
  }

  async function getValidationsResult() {
    const accountId = getCookie('account_id') || ''
    if (!accountId) {
      console.log('No account ID cookie found.')
    }

    const params: Record<string, string> = {
      show_details: 'true',
      account_id: accountId
    }

    try {
      const query = new URLSearchParams(params).toString()
      const response = await apiValidate(`/result?${query}`, {
        method: 'GET',
        upload: false
      })
      // console.log(response);

      return response
    } catch (error) {
      console.error('Error getting the validation result:', error)
      throw error
    }
  }

  function getCookie(name: string): string | null {
    const value = `; ${document.cookie}`
    const parts = value.split(`; ${name}=`)
    if (parts.length === 2) {
      const cookieValue = parts.pop()
      return cookieValue ? cookieValue.split(';').shift() || null : null
    }
    return null
  }
  return { createValidation, uploadImages, getValidationsResult, getInfoCountries }
}
