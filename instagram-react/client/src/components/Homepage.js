import React from "react";
import "../body.css";
import "bootstrap/dist/css/bootstrap.min.css";
import Header from "./Header";
import Content from "./Content";
class Homepage extends React.Component {
  render() {
    return (
      <div>
        <Header />
        <Content />
      </div>
    );
  }
}

export default Homepage;
