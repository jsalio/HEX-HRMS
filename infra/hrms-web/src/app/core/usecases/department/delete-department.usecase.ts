import { Inject, Injectable } from "@angular/core";
import { DEPARTMENT_REPOSITORY, DepartmentRepository } from "../../domain/ports/department.repo";

@Injectable({ providedIn: 'root' })
export class DeleteDepartmentUseCase {
    constructor(@Inject(DEPARTMENT_REPOSITORY) private departmentRepo: DepartmentRepository) {}
    
    async Execute(id: string): Promise<void> {
        return this.departmentRepo.delete(id);
    }
}
