import type {
  IGetImagesRequest,
  IGetImagesResponse,
  TImages,
  TOrderBy,
  TOrderDirection
} from '../../definitions'

import {
  EApiRoute
} from '../../definitions'

import api from '../api'

/**
 * Lekérdezi a képeket megadott rendezés szerint.
 * @param orderBy     - Miszerint legyen rendezve.
 * @param direction   - Milyen irányban.
 */
export default async function get (
  orderBy: TOrderBy = 'name',
  direction: TOrderDirection = 'asc'
): Promise<TImages> {
  const params: IGetImagesRequest = { direction, orderBy }

  const response = await api.get<IGetImagesResponse>(EApiRoute.Get, params)

  return response.isSuccess ? response.data.images : []
}