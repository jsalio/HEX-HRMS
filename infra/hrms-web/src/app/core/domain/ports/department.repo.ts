import { InjectionToken } from "@angular/core";
import { PaginatedResponse, SearchQuery, Department } from "../models";

export interface DepartmentRepository {
    list(query: SearchQuery): Promise<PaginatedResponse<Department>>;
    create(department: Department): Promise<Department>;
    update(department: Department): Promise<Department>;
    delete(id: string): Promise<void>;
    get(id: string): Promise<Department>;
}

export const DEPARTMENT_REPOSITORY = new InjectionToken<DepartmentRepository>('department.repository');
