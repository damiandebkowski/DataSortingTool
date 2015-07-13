var React = require('react');
var ReactAddon = require('react/addons');
var Bootstrap = require('bootstrap');
var _ = require('lodash');
var $ = require('jquery');

var modelObject = {
    chart: {
        plotBackgroundColor: null,
        plotBorderWidth: null,
        plotShadow: false,
        type: 'pie'
    },
    title: {
        text: 'Title Example'
    },
    tooltip:{
        pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
    },
    plotOptions: {
        pie: {
            allowPointSelect: true,
            cursor: 'pointer',
            dataLabels: {
                enabled: false
            },
            showInLegend: true
        }
    }
}

var seriesObject = [{
    name: 'Data Type',
    colorByPoint: true,
    data: [
        {
            name: 'One',
            y: 50.00
        },
        {
            name: 'Two',
            y: 50.00
        }
    ]
}];


var PieChart = React.createClass({
    componentDidMount: function(){
        var reference;
        switch(this.props.refName){
            case 'genderChart':
                reference = this.refs.genderChart;
                _.set(modelObject, 'title.text', 'Gender');
                //_.set(seriesObject[0], 'data[0].name', 'Male');
                //_.set(seriesObject[0], 'data[1].name', 'Female');
                break;
            case 'ageChart':
                reference = this.refs.ageChart;
                _.set(modelObject, 'title.text', 'Age');
                break;
            case 'locationChart':
                reference = this.refs.locationChart;
                _.set(modelObject, 'title.text', 'Location');
                break;
            case 'educationChart':
                reference = this.refs.educationChart;
                _.set(modelObject, 'title.text', 'Education');
                break;
        }
        var chartModel = _.cloneDeep(modelObject);
        var seriesModel = _.cloneDeep(seriesObject);
        var chartOptions = ReactAddon.addons.update(chartModel, {chart: {renderTo: {$set: reference.getDOMNode()}}, series: {$set: seriesModel}});
        var chartInstance = new window.Highcharts.Chart(chartOptions);
        this.setState({
            chartInstance: chartInstance
        });
    },
    updateGraph: function(){
        $.get("/GraphData", {type: this.props.refName}.bind(this), function(data){}.bind(this), "text");
    },
    render: function(){
        return(
            <div ref={this.props.refName}/>
        );
    }
});

module.exports = PieChart;
