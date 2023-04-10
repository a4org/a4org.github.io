/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./src/bond.js":
/*!*********************!*\
  !*** ./src/bond.js ***!
  \*********************/
/***/ (() => {

eval("String.prototype.capitalize = function() {\n    return this.replace(/(?:^|\\s)\\S/g, function(a) { return a.toUpperCase(); });\n};\n\nfunction figure3() {\n    var margin = ({\n        top: 30,\n        right: 30,\n        bottom: 30,\n        left: 60\n    });\n    \n    var indicator_image_size = 75;\n    var indicator_image_padding = 10;\n    var indicator_box_top_padding = 25;\n    \n    var tooltip_image_size = 100;\n\n    var chart_width = 650;\n    var chart_height = 300;\n    var chart_padding = 80;\n    \n    var width = chart_width;\n    var height = chart_height + indicator_box_top_padding + indicator_image_padding + indicator_image_size + chart_padding;\n    \n    var base_image_name = '美國國債孳息曲線(10-2年)';\n    var current_data = null;\n\n    var bar_colors = {\n        '美國國債孳息曲線(10-2年)': 'rgb(233, 205, 73)',\n        'JPM環球高風險債利差': 'rgb(211, 86, 42)',\n    }\n\n    var indicator_data = [\n        { x: 0, y: 0, id: '美國國債孳息曲線(10-2年)', opacity: 1.0},\n        { x: indicator_image_size + indicator_image_padding, y: 0, id: 'JPM環球高風險債利差', opacity: 0.2 },\n    ]\n    var container = d3.select('#bond')\n                        .append('svg')\n                        .attr('width',  '100%')\n                        .attr('height', '100%')\n                        .style('min-width', `${(width + margin.left + margin.right ) / 2}px`)\n                        .style('max-width', `${width + margin.left + margin.right}px`)\n                        .attr('viewBox', `0 0 ${width + margin.left + margin.right} ${height + margin.top + margin.bottom}`);\n    \n    var indicator_group = container\n        .append('g')\n        .attr('id', 'indicator_group')\n        .attr('width', 4 * indicator_image_size + 3 * indicator_image_padding)\n        .attr('height', indicator_image_size + indicator_image_padding)\n        .attr('transform', `translate(${margin.left}, ${margin.top + chart_height + indicator_box_top_padding + chart_padding})`);\n    indicator_group\n        .append('text')\n        .attr('x', indicator_group.attr('width') / 2)\n        .attr('y', -indicator_box_top_padding / 2)\n        .attr('text-anchor', 'middle')\n        .style('font-weight', 700)\n        .style('font-size', '18px')\n        .text('選擇債卷名稱:')\n            \n    var chart_group = container\n        .append('g')\n        .attr('id', 'chart_group')\n        .attr('width', chart_width)\n        .attr('height', chart_height)\n        .attr('transform', `translate(${margin.left}, ${margin.top})`);\n    \n    chart_group.append('text')\n        .attr(\"transform\", \"rotate(-90)\")\n        .attr(\"y\", 0 - chart_padding / 2 - 10)\n        .attr(\"x\", 0 - chart_height / 2)\n        .attr(\"dy\", \"1em\")\n        .style(\"text-anchor\", \"middle\")\n        \n    chart_group.append(\"text\")             \n      .attr(\"transform\", `translate(${(chart_width / 2)}, ${(chart_height + chart_padding / 2)})`)\n      .style(\"text-anchor\", \"middle\")\n\n    chart_group.append(\"text\")\n      .attr('id', 'chart_title')\n      .attr(\"transform\", `translate(${(chart_width) / 2}, -10)`)\n      .style(\"text-anchor\", \"middle\")\n      .style(\"font-weight\", 700)\n      .text(\"債卷名稱： \" + base_image_name.replace(/_/g, ' ').capitalize());\n    \n    container.selectAll('text').style(\"font-family\", \"sans-serif\");\n    \n    function draw_chart(data, should_init) {\n        current_data = data;\n        var x = d3.scaleTime()\n            .range([0, chart_width])\n            .domain(d3.extent(current_data, function(d) { return new Date(d.Date); }));\n\n        var y = d3.scaleLinear()\n        .range([chart_height, 0])\n        .domain([d3.min(current_data, function(d) { return Math.min(+d['美國國債孳息曲線(10-2年)'], +d['JPM環球高風險債利差']); }),\n          d3.max(current_data, function(d) { return Math.max(+d['美國國債孳息曲線(10-2年)'], +d['JPM環球高風險債利差']); })]);\n            \n        var line = d3.line()\n            .x(function(d) { return x(new Date(d.Date)) }) // Use date for x-axis\n            .y(function(d) { return y(+d[[base_image_name]])})\n            .curve(d3.curveMonotoneX);\n    \n        \n        if (should_init) {\n            var xaxis = chart_group.append('g')\n                .attr('class', 'axis axis--x')\n                .attr('transform', `translate(0, ${chart_height})`)\n                .attr('id', 'chart-x-axis')\n                .call(d3.axisBottom(x).ticks(d3.timeMonth.every(1)).tickFormat(d3.timeFormat('%m/%Y')));\n            \n            var yaxis = chart_group.append('g')\n                .attr('class', 'axis axis--y')\n                .attr('id', 'chart-y-axis')\n                .call(d3.axisLeft(y));\n            \n            chart_group.append('g')\n                .attr('id', 'grid_markings_horz')    \n                .selectAll('line.horizontalGrid').data(y.ticks()).enter()\n                .append('line')\n                .attr('class', 'horizontalGrid')\n                .attr('x1', 0)\n                .attr('x2', chart_width)\n                .attr('y1', function(d) { return y(d) + 0.5; })\n                .attr('y2', function(d) { return y(d) + 0.5; })\n                .attr('shape-rendering', 'crispEdges')\n                .attr('fill', 'none')\n                .attr('stroke', 'gray')\n                .attr('stroke-width', '1px')\n                .attr('stroke-opacity', 0.3);\n\n          chart_group.append('g')\n              .attr('id', 'grid_markings_vert')    \n              .selectAll('line.verticalGrid').data(x.ticks()).enter()\n              .append('line')\n              .attr('class', 'verticalGrid')\n              .attr('x1', function(d) { return x(d) + 0.5; })\n              .attr('x2', function(d) { return x(d) + 0.5; })\n              .attr('y1', 0)\n              .attr('y2', chart_height)\n              .attr('shape-rendering', 'crispEdges')\n              .attr('fill', 'none')\n              .attr('stroke', 'gray')\n              .attr('stroke-width', '1px')\n              .attr('stroke-opacity', 0.3);\n                \n            var path = chart_group.append('path')\n                .datum(current_data)\n                .attr('id', 'line_mark')\n                .attr('class', 'line')\n                .attr('d', line)\n                .attr('fill', 'none')\n                .attr('stroke', bar_colors[[base_image_name]])\n                .attr('stroke-width', 4);\n            \n            var w = tooltip_image_size - 10;\n            var h = tooltip_image_size - 10;\n\nvar tooltip = d3.select(\"body\").append(\"div\")\n    .attr(\"class\", \"tooltip\")\n    .style(\"opacity\", 0)\n    .style(\"position\", \"absolute\");\n\nchart_group.append(\"rect\")\n    .attr(\"fill\", \"none\")\n    .attr(\"pointer-events\", \"all\")\n    .attr(\"width\", chart_width)\n    .attr(\"height\", chart_height)\n    .on(\"mouseover\", function() { tooltip.style(\"opacity\", 1); })\n    .on(\"mouseout\", function() { tooltip.style(\"opacity\", 0); })\n    .on(\"mousemove\", mousemove);\n\nvar bisectDate = d3.bisector(function(d) { return new Date(d.Date); }).left;\n\nfunction mousemove() {\n        return;\n  var x0 = x.invert(d3.mouse(this)[0]);\n  var i = bisectDate(current_data, x0, 1);\n  var d0 = current_data[i - 1];\n  var d1 = current_data[i];\n\n  if (d0 === undefined || d1 === undefined) {\n      return;\n  }\n\n  var d = x0 - new Date(d0.Date) > new Date(d1.Date) - x0 ? d1 : d0;\n\n  tooltip.html(\"Date: \" + d3.timeFormat(\"%Y-%m-%d\")(new Date(d.Date)) + \"<br/>Price: \" + d[base_image_name])\n      .style(\"left\", (d3.event.pageX + 10) + \"px\")\n      .style(\"top\", (d3.event.pageY - 28) + \"px\")\n      .style(\"opacity\", 1);\n}\n\n\n        }\n        else {\n            var transition_duration = 500;\n            var mark_transition = d3\n                .transition()\n                .duration(transition_duration);\n            chart_group.select('#line_mark')\n                .datum(current_data)\n                .transition(mark_transition)\n                .attr('d', line)\n                .attr('stroke', bar_colors[[base_image_name]])\n            chart_group.select('#chart-x-axis')\n                .transition(mark_transition)\n                .call(d3.axisBottom(x).ticks(d3.timeMonth.every(1)).tickFormat(d3.timeFormat('%m/%Y')));\n        }\n    }\n    \n    function select_new_image(row, i) {\n        if (base_image_name === row.id) {\n            return;\n        }\n        \n        var indicator_images = indicator_group.selectAll('image').data(indicator_data)\n        indicator_images.attr('opacity', function(d) {\n            if (row.id == d.id) {\n                return 1.0;\n            } else {\n                return 0.2\n            }\n        })\n        \n        base_image_name = row.id;\n        chart_group.select('#chart_title')\n            .text(\"債卷名稱: \" + base_image_name.replace(/_/g, ' ').capitalize());\n        draw_chart(current_data, false);\n    }\n    \n    var indicator_images = indicator_group.selectAll('image').data(indicator_data);\n    indicator_images.enter()\n        .append('image')\n        .attr('width', indicator_image_size)\n        .attr('height', indicator_image_size)\n        .attr('xlink:href', function(d) {\n            return 'images/jpm.webp';\n        })\n        .attr('id', function(d) { return d.id; })\n        .attr('x', function(d) { return d.x; })\n        .attr('y', function(d) { return d.y; })\n        .attr('opacity', function(d) { return d.opacity; })\n        .on('click', select_new_image);\n    \n    d3.csv('data/macroindicators.csv').then(function(data) { \n      data.forEach(function(d) {\n          d.Date = d3.timeParse('%m/%d/%Y')(d.Date);\n      });\n      draw_chart(data, true) \n    });\n}\n\nfigure3();\n\n\n//# sourceURL=webpack://radar-index/./src/bond.js?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = {};
/******/ 	__webpack_modules__["./src/bond.js"]();
/******/ 	
/******/ })()
;