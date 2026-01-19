import { Component, OnInit, signal } from '@angular/core';
import { ListDepartmentUseCase } from '../../../core/usecases/department/list-department.usecase';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  departmentCount = signal<number>(0);

  constructor(private listDepartmentUseCase: ListDepartmentUseCase) {}

  async ngOnInit(): Promise<void> {
    try {
      const response = await this.listDepartmentUseCase.Execute({
        filters: [],
        pagination: { page: 1, limit: 1 }
      });
      this.departmentCount.set(response.total_rows);
    } catch (error) {
      console.error('Error fetching department count:', error);
    }
  }
}
