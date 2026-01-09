import { Department } from './department.model';

/**
 * Position model
 */
export interface Position {
  /** ID of the position */
  id: string;
  /** Name of the position */
  name: string;
  /** ID of the department */
  departmentId: string;
  /** Department of the position */
  department?: Department;
}
