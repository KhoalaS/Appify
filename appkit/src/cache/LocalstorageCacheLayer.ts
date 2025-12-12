import { CacheLayer } from './CacheLayer'

export class LocalstorageCacheLayer implements CacheLayer {
  async get<T>(key: string): Promise<T> {
    const rawValue = localStorage.getItem(key)
    try {
      return JSON.parse(rawValue)
    } catch {
      return rawValue as T
    }
  }

  async put<T>(key: string, value: T): Promise<void> {
    localStorage.setItem(key, JSON.stringify(value))
  }

  async delete(key: string): Promise<void> {
    localStorage.removeItem(key)
  }
}
