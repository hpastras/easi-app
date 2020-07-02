import React from 'react';
import { Link } from 'react-router-dom';
import Button from 'components/shared/Button';
import './index.scss';

const SideNavActions = () => {
  return (
    <div className="sidenav-actions grid-row flex-column">
      <div className="grid-col margin-top-105">
        <Link to="/">Save & Exit</Link>
      </div>
      <div className="grid-col margin-top-2">
        {/* Leaving this a button as it will likely do more than redirect the user */}
        <Button type="button" unstyled onClick={() => {}}>
          Remove your request to add a new system
        </Button>
      </div>
      <div className="grid-col margin-top-5">
        <h4>Related Content</h4>
        <a href="/governance-overview" target="_blank">
          Overview for adding a system
        </a>
        <p>(opens in a new tab)</p>
      </div>
    </div>
  );
};

export default SideNavActions;