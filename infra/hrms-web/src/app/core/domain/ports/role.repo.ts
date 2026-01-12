import { InjectionToken } from "@angular/core";
import { PaginatedResponse, SearchQuery } from "../models";
import { Role } from "../models/role.model";
import { Permission } from "../models/permission.model";

export interface RoleRepository {
    list(query: SearchQuery): Promise<PaginatedResponse<Role>>;
    create(role: Role): Promise<Role>;
    update(role: Role): Promise<Role>;
    get(id: string): Promise<Role>;
    getAll(query: SearchQuery): Promise<PaginatedResponse<Role>>;
    listSystemPermissions(): Promise<Permission[]>;
}

export const ROLE_REPOSITORY = new InjectionToken<RoleRepository>('role.repository');
