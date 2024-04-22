export interface IImage {
  readonly name: string
  readonly createdAt: string
}

export type TImages = readonly IImage[]

export type TOrderBy = 'name' | 'date'

export type TOrderDirection = 'asc' | 'desc'

export interface IGetImagesRequest {
  readonly orderBy: TOrderBy
  readonly direction: TOrderDirection
}

export interface IGetImagesResponse {
  readonly images: TImages
}

export enum EApiRoute {
  DeleteByName  = '/api/images/:name',
  Get           = '/api/images',
  Insert        = '/api/images'
}