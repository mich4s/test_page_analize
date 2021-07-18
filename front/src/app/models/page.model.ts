export interface PageModel {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  URL: string;
  Title: string;
  HTMLVersion: string;
  HeadingsCount: number;
  internalLinksCount: number;
  externalLinksCount: number;
  hasLoginForm: boolean;
}
