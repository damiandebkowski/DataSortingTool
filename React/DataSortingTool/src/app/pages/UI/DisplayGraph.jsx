var React = require('react');
var ReactAddon = require('react/addons');
var Bootstrap = require('bootstrap');
var _ = require('lodash');
var PieChart = require('./Charts/PieChart.jsx');
var $ = require('jquery');

/*
 * Class: DisplayGraph
 * Usage:
 * Design:
 */

var DisplayGraph = React.createClass({
    getInitialState: function(){
        return{active: this.props.dataTypes}
    },
    componentWillReceiveProps: function(nextProps){
        this.setState({active: nextProps.dataTypes});
    },
    shouldComponentUpdate: function(nextProps, nextState){
        if(this.props.buttonID != nextProps.buttonID){
            console.log("Go Get Server Data")
            return true;
        }
        return false;
    },
    render: function(){
        var genderChart;
        var ageChart;
        var locationChart;
        var educationChart;

        var findDataType = this.state.active;
        findDataType = findDataType.toString();

        if(findDataType.search("Gender") != -1){
            genderChart = <div className="col-sm-3"><PieChart refName="genderChart"/></div>
        }
        if(findDataType.search("Age") != -1){
            ageChart = <div className="col-sm-3"><PieChart refName="ageChart"/></div>
        }
        if(findDataType.search("Location") != -1){
            locationChart = <div className="col-sm-3"><PieChart refName="locationChart"/></div>
        }
        if(findDataType.search("Education") != -1){
             educationChart = <div className="col-sm-3"><PieChart refName="educationChart"/></div>
        }

        var errorType;
        if(findDataType.length == 0){
            errorType = <h3 className="text-center"><span className="label label-danger">Please Select A Data Type</span></h3>
        }

        return(
            <div className="panel panel-default">
                <h3 className="text-center">
                    <span className="label label-success">Analyzed Data:</span>
                </h3>
                {errorType}
                <div className="panel-body">
                    {genderChart}
                    {ageChart}
                    {locationChart}
                    {educationChart}
                </div>
            </div>
        );
    }
});

module.exports = DisplayGraph;
