interface AndroidBridge {
    showToast(message: string): void;
    saveFile(filename: string, content: string): void;
}

declare global {
    const Android: AndroidBridge;
}

export {};
