export class DynamicResult<T = any> {
  constructor(private data: T) {}

  getData(): T {
    return this.data;
  }

  convert<U>() {
    return {}
     //this.data as U;
  }

  static create<T>(data: T): DynamicResult<T> {
    return new DynamicResult(data);
  }
}
