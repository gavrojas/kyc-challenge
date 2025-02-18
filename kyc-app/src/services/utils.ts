import type { OptionsApiCall } from '@/types/index'

const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/validations'

type ApiCallOptions = OptionsApiCall | undefined

export async function apiValidate(
  path: string,
  { method = 'GET', data, headers, upload }: ApiCallOptions = {}
) {
  // const accountSession = sessionStorage.getItem('account-id')
  const requestHeaders: { [key: string]: string } = {
    ...headers
  }

  let body: BodyInit | undefined

  if (upload === true) {
    body = data
  } else {
    // Si no es FormData, asume que es un objeto JSON
    body = method !== 'GET' && data ? JSON.stringify(data) : undefined
    requestHeaders['Content-Type'] = 'application/json'
  }

  try {
    const response = await fetch(BASE_URL + path, {
      method: method,
      headers: requestHeaders,
      body
    })

    if (!response.ok) {
      const errorData = await response.json()
      if (response.status === 403) {
        // colocar lógica cierre sesión y eliminar registro cuenta en la db**
        throw new Error('Session expired. Recharging')
      }
      throw new Error(errorData.error || 'Request failed')
    }

    return response.json()
  } catch (error) {
    throw new Error(`API call failed: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}
