import { EApiRoute } from '../../definitions'

import api from '../api'

/**
 * Egy új kép beszúrása.
 * @param data - A küldendő formData.
 */
export default async function insert (data: FormData): Promise<boolean> {
  const res = await api.form(EApiRoute.Insert, data)

  return res.isSuccess
}