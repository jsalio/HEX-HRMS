import { Department } from './department.model';
import { User } from './user.model';

/**
 * Work type enum
 */
export type WorkType = 'Remote' | 'Hybrid' | 'OnSite';

/**
 * Position status enum
 */
export type PositionStatus = 'Active' | 'Inactive' | 'Closed';

/**
 * Position model representing a job position
 */
export interface Position {
  /** Unique identifier */
  id: string;
  /** Position title */
  title: string;
  /** Position code (unique) */
  code: string;
  /** Description of the position */
  description: string;
  /** Required skills (text) */
  requiredSkills: string;
  /** Minimum salary */
  salaryMin: number;
  /** Maximum salary */
  salaryMax: number;
  /** Maximum number of employees to hire */
  maxEmployees: number;
  /** Currency code */
  currency: string;
  /** Work type (Remote, Hybrid, OnSite) */
  workType: WorkType;
  /** Department ID */
  departmentId: string;
  /** Associated department */
  department?: Department;
  /** Status (Active, Inactive, Closed) */
  status: PositionStatus;
  /** Creation timestamp */
  createdAt?: string;
  /** Last update timestamp */
  updatedAt?: string;
  /** User who created this position */
  createdById?: number;
  /** User who last updated this position */
  updatedById?: number;
  /** User reference */
  user?: User;
}

/**
 * DTO for creating a new position
 */
export interface CreatePositionDto {
  title: string;
  code: string;
  description: string;
  requiredSkills: string;
  salaryMin: number;
  salaryMax: number;
  maxEmployees: number;
  currency: string;
  workType: WorkType;
  departmentId: string;
  createdById?: number;
}

/**
 * DTO for updating an existing position
 */
export interface UpdatePositionDto {
  id: string;
  title: string;
  code: string;
  description: string;
  requiredSkills: string;
  salaryMin: number;
  salaryMax: number;
  maxEmployees: number;
  currency: string;
  workType: WorkType;
  departmentId: string;
  status: PositionStatus;
  updatedById?: number;
}
