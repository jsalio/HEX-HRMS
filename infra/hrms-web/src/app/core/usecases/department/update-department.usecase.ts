import { Inject, Injectable } from "@angular/core";
import { Department } from "../../domain/models";
import { DEPARTMENT_REPOSITORY, DepartmentRepository } from "../../domain/ports/department.repo";

@Injectable({ providedIn: 'root' })
export class UpdateDepartmentUseCase {
    constructor(@Inject(DEPARTMENT_REPOSITORY) private departmentRepo: DepartmentRepository) {}
    
    async Execute(department: Department): Promise<Department> {
        return this.departmentRepo.update(department);
    }
}
