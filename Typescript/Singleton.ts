export class Singleton {
    private static instance: Singleton;
    private _data: number;
    private constructor() { }
    static getInstance() {
        if (!Singleton.instance) {
            Singleton.instance = new Singleton();
            Singleton.instance._data = 0;
        }
        return Singleton.instance;
    }
    get data(): number {
        return this._data;
    }
    set data(data) {
        this._data = data;
    }
}