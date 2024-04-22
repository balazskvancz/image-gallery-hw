import { EApiRoute } from '../../definitions'

import api from '../api'

/**
 * Töröl egy fájlt név szerint.
 * @param name - A törlendő egyed neve.
 */
export default async function deleteByName (
  name: string
): Promise<boolean> {
  const url = EApiRoute.DeleteByName.replace(":name", name)

  const { isSuccess } = await api.delete(url)

  return isSuccess
}