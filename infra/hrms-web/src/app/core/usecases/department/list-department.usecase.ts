import { Inject, Injectable } from "@angular/core";
import { PaginatedResponse, SearchQuery, Department } from "../../domain/models";
import { DEPARTMENT_REPOSITORY, DepartmentRepository } from "../../domain/ports/department.repo";

@Injectable({ providedIn: 'root' })
export class ListDepartmentUseCase {
    constructor(@Inject(DEPARTMENT_REPOSITORY) private departmentRepo: DepartmentRepository) {}
    
    async Execute(query: SearchQuery): Promise<PaginatedResponse<Department>> {
        return this.departmentRepo.list(query);
    }
}
