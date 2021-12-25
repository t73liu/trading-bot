import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";

import { preventDefault } from "../../utils/function";
import Title from "../common/Title";

const Deposits = () => (
  <>
    <Title>Recent Deposits</Title>
    <Typography component="p" variant="h4">
      $3,024.00
    </Typography>
    <Typography color="text.secondary" sx={{ flex: 1 }}>
      on 15 March, 2019
    </Typography>
    <div>
      <Link
        component="button"
        color="primary"
        href="/balances"
        onClick={preventDefault}
      >
        View balance
      </Link>
    </div>
  </>
);
export default Deposits;
