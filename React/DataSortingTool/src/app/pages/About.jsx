var React = require('react');
var Bootstrap = require('bootstrap');

var About = React.createClass({
    render: function(){
        return(
            <div>
                <div className="row">
                    <div className="col-md-12">
                        <h1 className="text-center">Welcome To DataSortingTool</h1>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-offset-1">
                        <h3><strong>Objective:</strong></h3>
                    </div>
                    <div className="col-md-10 col-md-offset-1">
                        <p className="text-justify">
                            My goal is to provide users on facebook with an idea of how their content affects their audience. By showing the user demographic
                            statistics through their previous posts, photos, videos, and statuses; the user can analyze how they can create content to target a specific audience range based on
                            age, gender, education, and location.
                        </p>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-offset-1">
                        <h3><strong>How?</strong></h3>
                    </div>
                    <div className="col-md-10 col-md-offset-1">
                        <p className="text-justify">
                            By logging into facebook on my website, you grant my app special permissions from facebook to utilize information from you facebook account. My app is only granted
                            information from your posts, videos, photos, and statuses. I then can request access to specific information from facebook's servers based on those permissions.
                            With the information I obtain, I am able to display it through pi chart graphs to help the user digest what they are analyzing.
                        </p>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-offset-1">
                        <h3><strong>Instructions</strong></h3>
                    </div>
                    <div className="col-md-10 col-md-offset-1">
                        <ol>
                            <li>Log Into Facebook By Clicking The Green 'Facebook Login' Button In The Upper Right Corner Of The Website</li>
                            <li>Make Sure You Are On The Main Page To Access The App, Which Can Be Accessed When Clicking The 'Home' Button Near The Upper Left Corner Of The Website</li>
                            <li>Select A Data Type To Analyze In The 'Select Data Types:' Panel</li>
                            <li>Click On One Of The Buttons: 'Posts, Pictures, Videos, or Status' To Load A List Of Your Content Within The Panel</li>
                            <li>Select One Of Your Contents From Facebook To Analyze The Demographic Statiscs</li>
                            <li>Scroll Down To View The Information In the Pie Charts</li>
                        </ol>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = About;
