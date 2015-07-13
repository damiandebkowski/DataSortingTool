var React = require('react');
var Bootstrap = require('bootstrap');

var Contact = React.createClass({
    render: function(){
        return(
            <div>
                <div className="row">
                    <div className="col-md-12">
                        <h1 className="text-center">Contact</h1>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-12">
                        <h4 className="text-center"><strong>Incase Of Any Bugs In The App You May Find, Please Use The Following Information To Contact Me. Thank You For Your Support!</strong></h4>i
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-11 col-md-offset-1">
                        <p className="text-justify">
                            Email: Damian3395@gmail.com
                        </p>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-1 col-md-offset-1">
                        <p>GitHub:</p>
                    </div>
                    <div className="col-md-1">
                        <button type="button" className="btn btn-link">github.com/damiandebkowski/DataSortingTool/issues</button>
                    </div>
                </div>
            </div>
        );
    }
});

module.exports = Contact;
