export interface CacheLayer {
  get<T>(key: string): Promise<T>
  put<T>(key: string, value: T): Promise<void>
  delete(key: string): Promise<void>
}
