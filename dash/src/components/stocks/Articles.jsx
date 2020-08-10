import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@material-ui/core";
import PropTypes from "prop-types";
import dayjs from "dayjs";
import ExternalLink from "../common/ExternalLink";
import OverflowDiv from "../common/OverflowDiv";

const Articles = ({ articles }) => (
  <TableContainer>
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Title</TableCell>
          <TableCell>Description</TableCell>
          <TableCell>Published At</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {articles?.map((article) => (
          <TableRow key={article.url}>
            <TableCell>
              <ExternalLink url={article.url}>{article.title}</ExternalLink>
            </TableCell>
            <TableCell>
              <OverflowDiv>{article.summary}...</OverflowDiv>
            </TableCell>
            <TableCell>
              {dayjs(article.publishedAt).format("YYYY-MM-DD HH:mm:ss")}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </TableContainer>
);

Articles.propTypes = {
  articles: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string.isRequired,
      summary: PropTypes.string.isRequired,
      url: PropTypes.string.isRequired,
      publishedAt: PropTypes.string.isRequired,
    })
  ).isRequired,
};

export default Articles;
