interface ImportMeta {
  readonly env: ImportMetaEnv;
}

interface ImportMetaEnv {
  /**
   * Internal variable. Do not use.
   */
  readonly [key: string]: any;

  readonly NG_APP_API_URL: string;
  // Add other variables here...
}
