export class IDBService {
  private database?: IDBDatabase
  private databaseVersion: number = 1
  private databaseName: string
  private registeredObjectStores = new Map<string, string[]>()

  constructor(databaseName: string, databaseVersion: number) {
    this.databaseName = databaseName
    this.databaseVersion = databaseVersion
  }

  registerObjectStore(name: string, keys: string[]) {
    this.registeredObjectStores.set(name, keys)
  }

  private getDb(): Promise<IDBDatabase> {
    return new Promise((resolve, reject) => {
      if (this.database == undefined) {
        const request = indexedDB.open(this.databaseName, this.databaseVersion)
        request.onerror = () => {
          reject(request.error)
        }

        request.onupgradeneeded = () => {
          this.database = request.result

          for (const [key, val] of this.registeredObjectStores) {
            if (this.database.objectStoreNames.contains(key)) {
              continue
            }
            const objectstore = this.database.createObjectStore(key, {
              keyPath: 'id',
            })
            val.forEach((field) => {
              objectstore.createIndex(field, field, {
                unique: false,
              })
            })
          }
          resolve(this.database)
        }

        request.onsuccess = () => {
          this.database = request.result
          resolve(this.database)
        }
      } else {
        return resolve(this.database)
      }
    })
  }

  async put(storeName: string, id: string, value: unknown) {
    const db = await this.getDb()
    const transaction = db.transaction([storeName], 'readwrite')
    const objectstore = transaction.objectStore(storeName)
    objectstore.put({ id, value })
  }

  async get<T>(storeName: string, id: string): Promise<T> {
    const db = await this.getDb()
    const transaction = db.transaction([storeName], 'readonly')
    const objectstore = transaction.objectStore(storeName)

    const req = objectstore.get(id)

    return new Promise((resolve, reject) => {
      req.onerror = () => {
        reject(req.error)
      }

      req.onsuccess = () => {
        resolve(req.result.value)
      }
    })
  }

  async getAll<T>(storeName: string): Promise<T[]> {
    const db = await this.getDb()

    return new Promise((resolve, reject) => {
      const tx = db.transaction(storeName, 'readonly')
      const store = tx.objectStore(storeName)
      const req = store.getAll()

      req.onsuccess = () => resolve(req.result)
      req.onerror = () => reject(req.error)
    })
  }

  async delete(storeName: string, id: string): Promise<void> {
    const db = await this.getDb()
    const transaction = db.transaction([storeName], 'readwrite')
    const objectStore = transaction.objectStore(storeName)

    const req = objectStore.delete(id)

    return new Promise((resolve, reject) => {
      req.onerror = () => {
        reject(req.error)
      }

      req.onsuccess = () => {
        resolve()
      }
    })
  }
}
