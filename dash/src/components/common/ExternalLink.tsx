import { ReactNode } from "react";

interface ExternalLinkProps {
  url: string;
  children?: ReactNode | undefined;
}

const ExternalLink = ({ url, children }: ExternalLinkProps): JSX.Element => (
  <a href={url} target="_blank" rel="noopener noreferrer">
    {children}
  </a>
);

export default ExternalLink;
