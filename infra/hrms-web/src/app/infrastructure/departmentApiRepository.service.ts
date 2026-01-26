import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom } from "rxjs";
import { PaginatedResponse, SearchQuery, Department } from "../core/domain/models";
import { DepartmentRepository } from "../core/domain/ports/department.repo";
import { environment } from "../../environments/environment.development";

@Injectable({ providedIn: 'root' })
export class DepartmentApiRepository implements DepartmentRepository {

  constructor(private http: HttpClient) {}

  async list(query: SearchQuery): Promise<PaginatedResponse<Department>> {
    const dto = await firstValueFrom(
      this.http.post<PaginatedResponse<Department>>(`${environment.apiUrl}/department/get-all`, query)
    );
    return dto;
  }

  async create(department: Department): Promise<Department> {
    const dto = await firstValueFrom(
      this.http.post<Department>(`${environment.apiUrl}/department/create`, department)
    );
    return dto;
  }

  async update(department: Department): Promise<Department> {
    const dto = await firstValueFrom(
      this.http.post<Department>(`${environment.apiUrl}/department/update`, department)
    );
    return dto;
  }

  async delete(id: string): Promise<void> {
    await firstValueFrom(
      this.http.post(`${environment.apiUrl}/department/delete`, { id })
    );
  }

  async get(id: string): Promise<Department> {
    const dto = await firstValueFrom(
      this.http.post<Department>(`${environment.apiUrl}/department/get`, { id })
    );
    return dto;
  }
}
