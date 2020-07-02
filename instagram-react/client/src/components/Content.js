import React from "react";
import history from "../History";
const bodyScrollLock = require('body-scroll-lock');
const disableBodyScroll = bodyScrollLock.disableBodyScroll;
const enableBodyScroll = bodyScrollLock.enableBodyScroll;

class Content extends React.Component {
  objs = [
    {
      ig: require("../img/rand1.jpg"),
      name: "abcd",
      liked: false,
      saved: false,
      followed: false,
    },
    {
      ig: require("../img/rand2.jpg"),
      name: "abcd",
      liked: false,
      saved: false,
      followed: false,
    },
    {
      ig: require("../img/rand3.jpg"),
      name: "abcd",
      liked: false,
      saved: false,
      followed: false,
    },
    {
      ig: require("../img/rand4.jpg"),
      name: "abcd",
      liked: false,
      saved: false,
      followed: false,
    },
    {
      ig: require("../img/rand5.jpg"),
      name: "abcd",
      liked: false,
      saved: false,
      followed: false,
    },
  ];
  ob = ["/img/rand1.jpg"];

  handlelike(e) {
    if (this.objs[e].liked === false) this.objs[e].liked = true;
    else this.objs[e].liked = false;

    history.push("/instagram-replica/home/");
  }
  handlesave(e) {
    if (this.objs[e].saved === false) this.objs[e].saved = true;
    else this.objs[e].saved = false;

    history.push("/instagram-replica/home/");
  }

  handlefollow(e) {
    if (this.objs[e].followed === false) this.objs[e].followed = true;
    else this.objs[e].followed = false;

    history.push("/instagram-replica/home/");
  }
  showoption() {
    document.getElementById("options").style.display = "block";
    const targetElement = document.querySelector('#root');
    disableBodyScroll(targetElement);
  }
  hideoptions() {
    document.getElementById("options").style.display = "none";
    const targetElement = document.querySelector('#root');
    enableBodyScroll(targetElement);
  }
  render() {
    const objs = this.objs;
    return (
      <div>
        <div>
          <div className="story-show">
            {objs.map((obj) => (
                <img className="ico" src={obj.ig} alt="" />
            ))}
          </div>
          <div className="side-show">
            <div className="user-info">
              <img src={require("../img/rand2.jpg")} alt="" />
              <div className="name">
                <b>Best uh!</b>
                <br />
                Best uh..! Best uh..!
              </div>
            </div>
            <div className="suggesition">
              <table style={{ cellspacing: "30px" }}>
                <thead>
                  <td>Suggestions for you</td>
                  <td id="right">See All</td>
                </thead>
                <br />
                {objs.map((obj, index) => (
                  <tbody>
                    <tr id="roow" key={index}>
                      <td>
                        <img className="icooo" src={obj.ig} alt="" />
                        <div className="name">
                          <b>{obj.name}</b>
                          <br />
                          <p>Follows you</p>
                        </div>
                      </td>
                      <td id="right">
                        <button
                          id="follow-button"
                          type="button"
                          onClick={(e) => {
                            this.handlefollow(index);
                          }}
                          style={{ border: "0", margin: "0", padding0: "10" }}
                        >
                          {obj.followed ? (
                            <p id="side-follow" style={{ color: "gray" }}>
                              Requested
                            </p>
                          ) : (
                            <p
                              id="side-follow"
                              style={{
                                color: "blue",
                                textDecoration: "underline",
                              }}
                            >
                              Follow
                            </p>
                          )}
                        </button>
                      </td>
                    </tr>
                  </tbody>
                ))}
              </table>
            </div>
            <div class="SkY6J">
              <nav class="uxKLF">
                <ul class="ixdEe _9Rlzb">
                  <li class="K5OFK">
                    <a
                      class="l93RR"
                      href="https://about.instagram.com/about-us"
                      rel="nofollow noopener noreferrer"
                      target="_blank"
                    >
                      About
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="https://help.instagram.com/">
                      Help
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="https://instagram-press.com/">
                      Press
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/developer/">
                      API
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/about/jobs/">
                      Jobs
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/legal/privacy/">
                      Privacy
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR _vfM2" href="/legal/terms/">
                      Terms
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/explore/locations/">
                      Locations
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/directory/profiles/">
                      Top Accounts
                    </a>
                  </li>
                  <li class="K5OFK">
                    <a class="l93RR" href="/directory/hashtags/">
                      Hashtags
                    </a>
                  </li>
                  <li class="K5OFK">
                    <span class="_3G4x7  l93RR">
                      Language
                      <select
                        aria-label="Switch Display Language"
                        class="hztqj"
                      >
                        <option value="en">English</option>

                        <option value="ta">தமிழ்</option>
                      </select>
                    </span>
                  </li>
                </ul>
              </nav>
              <span class="DINPA">© 2020 Instagram from Facebook</span>
            </div>
          </div>
        </div>
        {objs.map((obj, index) => (
          <div className="contents">
            <div className="top">
              <img className="ico" src={obj.ig} alt="" />
              <b>{obj.name}</b>
              <div class="MEAGs">
                <button class="wpO6b" type="button" onClick={this.showoption}>
                  <div
                    class="Igw0E rBNOH YBx95 _4EzTm "
                    style={{ height: "24px", width: "24px" }}
                  >
                    <svg
                      aria-label="More options"
                      class="_8-yf5 "
                      fill="#262626"
                      height="16"
                      viewBox="0 0 48 48"
                      width="16"
                    >
                      <circle
                        clip-rule="evenodd"
                        cx="8"
                        cy="24"
                        fill-rule="evenodd"
                        r="4.5"
                      ></circle>
                      <circle
                        clip-rule="evenodd"
                        cx="24"
                        cy="24"
                        fill-rule="evenodd"
                        r="4.5"
                      ></circle>
                      <circle
                        clip-rule="evenodd"
                        cx="40"
                        cy="24"
                        fill-rule="evenodd"
                        r="4.5"
                      ></circle>
                    </svg>
                  </div>
                </button>
              </div>
            </div>
            <img className="post" src={obj.ig} alt="" />
            <div className="bottom">
              <div className="ltpMr Slqrh">
                <span className="fr66n">
                  <button
                    className="wpO6b"
                    type="button"
                    onClick={(e) => {
                      this.handlelike(index);
                    }}
                  >
                    {obj.liked ? (
                      <svg
                        aria-label="Unlike"
                        class="_8-yf5 "
                        fill="#ed4956"
                        height="24"
                        viewBox="0 0 48 48"
                        width="24"
                      >
                        <path d="M34.6 3.1c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5s1.1-.2 1.6-.5c1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
                      </svg>
                    ) : (
                      <svg
                        aria-label="Like"
                        className="_8-yf5 "
                        fill="#262626"
                        height="24"
                        viewBox="0 0 48 48"
                        width="24"
                      >
                        <path d="M34.6 6.1c5.7 0 10.4 5.2 10.4 11.5 0 6.8-5.9 11-11.5 16S25 41.3 24 41.9c-1.1-.7-4.7-4-9.5-8.3-5.7-5-11.5-9.2-11.5-16C3 11.3 7.7 6.1 13.4 6.1c4.2 0 6.5 2 8.1 4.3 1.9 2.6 2.2 3.9 2.5 3.9.3 0 .6-1.3 2.5-3.9 1.6-2.3 3.9-4.3 8.1-4.3m0-3c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5.6 0 1.1-.2 1.6-.5 1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
                      </svg>
                    )}
                  </button>
                </span>
                <span className="_15y0l">
                  <button className="wpO6b " type="button">
                    <svg
                      aria-label="Comment"
                      className="_8-yf5 "
                      fill="#262626"
                      height="24"
                      viewBox="0 0 48 48"
                      width="24"
                    >
                      <path
                        clipRule="evenodd"
                        d="M47.5 46.1l-2.8-11c1.8-3.3 2.8-7.1 2.8-11.1C47.5 11 37 .5 24 .5S.5 11 .5 24 11 47.5 24 47.5c4 0 7.8-1 11.1-2.8l11 2.8c.8.2 1.6-.6 1.4-1.4zm-3-22.1c0 4-1 7-2.6 10-.2.4-.3.9-.2 1.4l2.1 8.4-8.3-2.1c-.5-.1-1-.1-1.4.2-1.8 1-5.2 2.6-10 2.6-11.4 0-20.6-9.2-20.6-20.5S12.7 3.5 24 3.5 44.5 12.7 44.5 24z"
                        fillRule="evenodd"
                      ></path>
                    </svg>
                  </button>
                </span>
                <button className="wpO6b " type="button">
                  <svg
                    aria-label="Share Post"
                    className="_8-yf5 "
                    fill="#262626"
                    height="24"
                    viewBox="0 0 48 48"
                    width="24"
                  >
                    <path d="M47.8 3.8c-.3-.5-.8-.8-1.3-.8h-45C.9 3.1.3 3.5.1 4S0 5.2.4 5.7l15.9 15.6 5.5 22.6c.1.6.6 1 1.2 1.1h.2c.5 0 1-.3 1.3-.7l23.2-39c.4-.4.4-1 .1-1.5zM5.2 6.1h35.5L18 18.7 5.2 6.1zm18.7 33.6l-4.4-18.4L42.4 8.6 23.9 39.7z"></path>
                  </svg>
                </button>
                <span className="wmtNn">
                  <button
                    className="wpO6b"
                    type="button"
                    onClick={(e) => {
                      this.handlesave(index);
                    }}
                  >
                    {obj.saved ? (
                      <svg
                        aria-label="Remove"
                        class="_8-yf5 "
                        fill="#262626"
                        height="24"
                        viewBox="0 0 48 48"
                        width="24"
                      >
                        <path d="M43.5 48c-.4 0-.8-.2-1.1-.4L24 28.9 5.6 47.6c-.4.4-1.1.6-1.6.3-.6-.2-1-.8-1-1.4v-45C3 .7 3.7 0 4.5 0h39c.8 0 1.5.7 1.5 1.5v45c0 .6-.4 1.2-.9 1.4-.2.1-.4.1-.6.1z"></path>
                      </svg>
                    ) : (
                      <svg
                        aria-label="Save"
                        className="_8-yf5 "
                        fill="#262626"
                        height="24"
                        viewBox="0 0 48 48"
                        width="24"
                      >
                        <path d="M43.5 48c-.4 0-.8-.2-1.1-.4L24 29 5.6 47.6c-.4.4-1.1.6-1.6.3-.6-.2-1-.8-1-1.4v-45C3 .7 3.7 0 4.5 0h39c.8 0 1.5.7 1.5 1.5v45c0 .6-.4 1.2-.9 1.4-.2.1-.4.1-.6.1zM24 26c.8 0 1.6.3 2.2.9l15.8 16V3H6v39.9l15.8-16c.6-.6 1.4-.9 2.2-.9z"></path>
                      </svg>
                    )}
                  </button>
                </span>
              </div>
              <div style={{ marginLeft: "6px" }}>
                <b>123 likes</b>
              </div>
            </div>
          </div>
        ))}
        <div id="options">
          <div class="mt3GC">
            <button class="aOOlW -Cab_ " tabindex="0">
              Report Inappropriate
            </button>
            <br />
            <button class="aOOlW -Cab_ " tabindex="0">
              Unfollow
            </button>
            <br />
            <button class="aOOlW   HoLwm " tabindex="0">
              Go to post
            </button>
            <br />
            <button class="aOOlW   HoLwm " tabindex="0">
              Share
            </button>
            <br />
            <button class="aOOlW   HoLwm " tabindex="0">
              Copy Link
            </button>
            <br />
            <button class="aOOlW   HoLwm " tabindex="0">
              Embed
            </button>
            <br />
            <button
              class="aOOlW   HoLwm "
              tabindex="0"
              onClick={this.hideoptions}
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    );
  }
}

export default Content;
