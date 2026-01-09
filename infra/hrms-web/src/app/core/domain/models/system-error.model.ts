export enum SystemErrorType {
  Internal = 'internal',
  Validation = 'validation'
}

export enum SystemErrorLevel {
  Info = 'info',
  Warning = 'warning',
  Error = 'error'
}

export enum SystemErrorCode {
  Internal = 500,
  Validation = 400,
  Migration = 404,
  None = 0
}

export interface SystemError {
  code: SystemErrorCode;
  type: SystemErrorType;
  level: SystemErrorLevel;
  message: string;
  details: Record<string, any>;
}

export function createSystemError(
  code: SystemErrorCode,
  type: SystemErrorType,
  level: SystemErrorLevel,
  message: string,
  details: Record<string, any> = {}
): SystemError {
  return {
    code,
    type,
    level,
    message,
    details
  };
}
