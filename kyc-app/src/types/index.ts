export interface OptionsApiCall {
  method?: string
  data?: any
  headers?: any
  upload?: boolean
}

export interface SessionData {
  account_id: string
  creation_date: Date
}

export type Country = 'Colombia' | 'Costa Rica' | 'Brasil' | 'Chile' | 'Peru' | 'Mexico' | 'Other'
export type DocType =
  | 'driver-license'
  | 'foreign-id'
  | 'identity-card'
  | 'national-id'
  | 'passport'
  | 'ppt'
  | 'rut'
  | 'cnh'
  | 'invoice'
  | 'picture-id'
  | 'record'

export type InfoCountries = {
  [key in Country]: {
    documentTypes: DocType[]
  }
}

export interface ValidationResultInfo {
  validation_status: string
  failure_status: string
  declined_reason: string
  processing_start_date: string
  processing_finish_date: string
  details: {
    document_details: DocumentDetails
  }
  user_response: {
    input_files: string[]
  }
}

export type DocumentDetails = {
  birth_place: string
  client_id: string
  country: string
  creation_date: string
  date_of_birth: string
  doc_id: string
  document_type: string
  document_version: string
  expedition_place: string
  gender: string
  height: string
  issue_date: string
  mime_type: string
  national_registrar: string
  production_data: string
  rh: string
  update_date: string
}
