import axios from 'axios'

type TAnyObject = Record<string, any>

type TMethod = 'get' | 'post' | 'delete'
type THeader = Record<string, string>

const axiosAPI = axios.create({
  baseURL:'http://localhost:3000', // TODO: mi legyen __PROD__ esetben.
  withCredentials: true,
  headers: {
    Accept: 'application/json'
  }
})

interface IResponse<T extends TAnyObject> {
  readonly isSuccess: boolean
  readonly data: T
}

// implement a method to execute all the request from here.
async function apiRequest<T extends TAnyObject = TAnyObject> (
  method: TMethod,
  url: string,
  data: TAnyObject | undefined | FormData,
  extraHeaders: THeader = {},
  params: TAnyObject = {}
): Promise<IResponse<T>> {
  const headers = {
    authorization: '',
    'Access-Control-Allow-Origin': '*',
    ...extraHeaders
  }

  try {
    const res = await axiosAPI({
      method,
      url,
      data,
      headers,
      params
    })

    return {
      isSuccess: true,
      data: res.data as T
    }
  }
  catch (err) {
    // TODO: handlerrror
    return {
      isSuccess: false,
      data: {} as T
    }
  }
}

/**
 * Get típusú kérés a megadott URL-re.
 * @param url           - Cél url.
 * @param extraHeaders  - Egyéb – explicit – fejlécek
 */
function get <T extends TAnyObject = TAnyObject> (
  url: string,
  params?: TAnyObject,
  extraHeaders?: THeader
): Promise<IResponse<T>> {
  return apiRequest('get', url, undefined, extraHeaders, params)
}

/**
 * Delete típusú kérés a megadott URL-re.
 * @param url           - Cél url.
 * @param extraHeaders  - Egyéb – explicit – fejlécek
 */
function deleteRequest <T extends TAnyObject = TAnyObject> (
  url: string,
  extraHeaders?: THeader
): Promise<IResponse<T>> {
  return apiRequest('delete', url, undefined, extraHeaders)
}

/**
 * Post típusú kérés a megadott URL-re.
 * @param url           - Cél url.
 * @param data          - A küldendő adat.
 * @param extraHeaders  - Egyéb – explicit – fejlécek
 */
function post <T extends TAnyObject = TAnyObject> (
  url: string,
  data: TAnyObject,
  extraHeaders?: THeader
): Promise<IResponse<T>> {
  const headers = {
    'Content-Type': 'application/json',
    ...extraHeaders
  }

  return apiRequest('post', url, data, headers)
}


/**
 * FormData típusú kérés küldése a megadott URL-re.
 * @param url           - Cél url.
 * @param data          - Küldendő formData.
 * @param extraHeaders  - Egyéb fejlécek.
 * @returns
 */
function form <T extends TAnyObject = TAnyObject> (
  url: string,
  data: FormData,
  extraHeaders?: THeader
): Promise<IResponse<T>> {
  const headers = {
    'Content-Type': 'multipart/form-data',
    ...extraHeaders
  }

  return apiRequest('post', url, data, headers)
}

// expose your method to other services or actions
const api = {
  get,
  delete: deleteRequest,
  post,
  form
}

export default api