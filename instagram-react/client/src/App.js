import React from "react";
import "./App.css";
import "./body.css";
import FloatingLabelInput from "react-floating-label-input";
import { Button } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import history from "./History";
import FacebookLogin from "react-facebook-login";

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <div className="logo"></div>
        <div className="form">
          <Logi />
        </div>
      </div>
    );
  }
}
class Logi extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
    };
    this.validateForm = this.validateForm.bind(this);
    this.redirect = this.redirect.bind(this);
  }

  validateForm() {
    return this.state.email.length > 0 && this.state.password.length > 6;
  }
  redirect() {
    history.push("/instagram-replica/home/");
  }

  render() {
    return (
      <div className="app">
        <div className="Login">
          <div className="email">
            <FloatingLabelInput
              id="email"
              type="text"
              label="Phone Number,Email,User name"
              value={this.state.email}
              onChange={(e) => {
                this.setState({ email: e.target.value });
              }}
            />
          </div>
          <div className="password">
            <FloatingLabelInput
              id="password"
              type="password"
              label="Password"
              value={this.state.password}
              onChange={(e) => {
                this.setState({ password: e.target.value });
              }}
            />
          </div>
          <Button block onClick={this.redirect} type="button">
            <b>Log in</b>
          </Button>
        </div>
        <div class="K-1uj Z7p_S"><div class="s311c"></div><div class="_0tv-g">or</div><div class="s311c"></div></div>
        <FacebookLogin
            appId="269025137482461"
            autoLoad={false}
            fields="name,email,picture"
            onClick={() => {
              console.log("click");
            }}
            callback={this.redirect}
            cssClass="facebook-button"
            icon="fa-facebook"

          />
        <div className="seperate">
          <br />
          <a href="/">Forgot password?</a>
        </div>
        <div className="signup">
          Don't have an account?
          <a href="/">
            <b>Sign up</b>
          </a>
        </div>
       
        
      </div>
    );
  }
}
export default App;



//downbar

/*
 <div class="APQi1">
           <p class="b_nGN">Get the app.</p>
           <div class="iNy2T">
             <a class="z1VUo" href="https://itunes.apple.com/app/instagram/id389801252?pt=428156&amp;ct=igweb.loginPage.badge&amp;mt=8&amp;vt=lo">
               <img alt="Available on the App Store" class="Rt8TI " src="/static/images/appstore-install-badges/badge_ios_english-en.png/180ae7a0bcf7.png"/></a>
             <a class="z1VUo" href="https://play.google.com/store/apps/details?id=com.instagram.android&amp;referrer=utm_source%3Dinstagramweb%26utm_campaign%3DloginPage%26ig_mid%3D917BFFEE-EC86-4711-86D4-119E373E0614%26utm_content%3Dlo%26utm_medium%3Dbadge">
            <img alt="Available on Google Play" class="Rt8TI " src="/static/images/appstore-install-badges/badge_android_english-en.png/e9cd846dc748.png"/></a>
           </div>
        </div>
*/