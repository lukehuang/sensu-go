import React from "react";
import PropTypes from "prop-types";
import map from "lodash/map";
import get from "lodash/get";
import { createFragmentContainer, graphql } from "react-relay";

import Table, {
  TableBody,
  TableHead,
  TableCell,
  TableRow,
} from "material-ui/Table";
import Checkbox from "material-ui/Checkbox";
import Row from "./CheckRow";

class CheckList extends React.Component {
  static propTypes = {
    viewer: PropTypes.shape({
      checks: PropTypes.shape({
        edges: PropTypes.array.isRequired,
      }),
    }).isRequired,
  };

  render() {
    const { viewer } = this.props;
    const checks = get(viewer, "checks.edges", []);

    return (
      <Table>
        <TableHead>
          <TableRow>
            <TableCell padding="checkbox">
              <Checkbox />
            </TableCell>
            <TableCell>Check</TableCell>
            <TableCell>Command</TableCell>
            <TableCell>Subscribers</TableCell>
            <TableCell>Interval</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {map(checks, edge => <Row key={edge.cursor} check={edge.node} />)}
        </TableBody>
      </Table>
    );
  }
}

export default createFragmentContainer(
  CheckList,
  graphql`
    fragment CheckList_viewer on Viewer {
      checks(first: 1500) {
        edges {
          node {
            ...CheckRow_check
          }
          cursor
        }
        pageInfo {
          hasNextPage
        }
      }
    }
  `,
);
