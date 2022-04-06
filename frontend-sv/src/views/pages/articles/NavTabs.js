import React from 'react'
import { CNav, CNavItem, CNavLink } from '@coreui/react'
import propTypes from 'prop-types'

function NavTabs(props) {
  return (
    <>
      <CNav variant="pills" role="tablist">
        <CNavItem>
          <CNavLink href="/all-post">Publish ({localStorage.getItem('publishLength')})</CNavLink>
        </CNavItem>
        <CNavItem>
          <CNavLink href="/draft">Draft ({localStorage.getItem('draftLength')})</CNavLink>
        </CNavItem>
        <CNavItem>
          <CNavLink href="/thrash">Trash ({localStorage.getItem('trashLength')})</CNavLink>
        </CNavItem>
      </CNav>
    </>
  )
}

NavTabs.propTypes = {
  draftLength: propTypes.number,
  thrashLength: propTypes.number,
  publishLength: propTypes.number,
}

export default NavTabs
