interface AndroidBridge {
  showToast(message: string): void
  saveFile(filename: string, content: string): void
  base64EncodeResponseToDataUrl(url: string): string
}

declare global {
  const Android: AndroidBridge
}

export {}
