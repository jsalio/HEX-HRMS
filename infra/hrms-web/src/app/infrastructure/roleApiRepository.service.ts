import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom } from "rxjs";
import { PaginatedResponse, SearchQuery, Role, Permission } from "../core/domain/models";
import { RoleRepository } from "../core/domain/ports/role.repo";
import { environment } from "../../environments/environment.development";

@Injectable({ providedIn: 'root' })
export class RoleApiRepository implements RoleRepository {

  constructor(private http: HttpClient) {}

  async list(query: SearchQuery): Promise<PaginatedResponse<Role>> {
    const dto = await firstValueFrom(
      this.http.post<PaginatedResponse<Role>>(`${environment.apiUrl}/roles/get-all`, query)
    );
    return dto;
  }

  async create(role: Role): Promise<Role> {
    const dto = await firstValueFrom(
      this.http.post<Role>(`${environment.apiUrl}/roles/create`, role)
    );
    return dto;
  }

  async update(role: Role): Promise<Role> {
    const dto = await firstValueFrom(
      this.http.post<Role>(`${environment.apiUrl}/roles/update`, role)
    );
    return dto;
  }

  async get(id: string): Promise<Role> {
     const dto = await firstValueFrom(
      this.http.post<Role>(`${environment.apiUrl}/roles/get`, { id })
    );
    return dto;
  }
  
  async getAll(query: SearchQuery): Promise<PaginatedResponse<Role>> {
    const dto = await firstValueFrom(
      this.http.post<PaginatedResponse<Role>>(`${environment.apiUrl}/roles/get-all`, query)
    );
    return dto;
  }

  async listSystemPermissions(): Promise<Permission[]> {
    const dto = await firstValueFrom(
      this.http.get<Permission[]>(`${environment.apiUrl}/roles/system-permissions`)
    );
    return dto;
  }

}
