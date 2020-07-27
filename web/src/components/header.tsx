import React from 'react';
import { AppBar, Toolbar, Typography, Button } from '@material-ui/core';

import { makeStyles, Theme } from '@material-ui/core/styles';

const useStyles = makeStyles((theme: Theme) => ({
  title: {
    fontSize: 30,
    fontWeight: 500,
    flexGrow: 1
  },
}));

const Header = (props: any) => {
  const styles = useStyles();
  let content = (
    <AppBar>
      <Toolbar>
        <Typography className={styles.title}>WantAPrice</Typography>
        <Button color="inherit">Log Out</Button>
      </Toolbar>
    </AppBar>
  );
  return content;
}

export default Header;