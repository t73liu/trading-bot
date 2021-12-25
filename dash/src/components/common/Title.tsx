import Typography from "@mui/material/Typography";
import { ReactNode } from "react";

interface TitleProps {
  children?: ReactNode;
}

const Title = ({ children }: TitleProps): JSX.Element => (
  <Typography component="h2" variant="h6" color="primary" gutterBottom>
    {children}
  </Typography>
);

export default Title;
