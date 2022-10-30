// Code generated by mdatagen. DO NOT EDIT.

package emareceiver

import (
	"fmt"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
	enabledProvidedByUser bool
	// Enabled               bool `mapstructure:"enabled"`
	// time_limit            string `mapstructure:"time_limit"`
	// limit            	  int `mapstructure:"limit"`
}

func (ms *MetricSettings) IsEnabledProvidedByUser() bool {
	return ms.enabledProvidedByUser
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}


// MetricsSettings provides settings for emareceiver metrics.
type MetricsSettings struct {
	MysqlLocks                 MetricSettings `mapstructure:"mysql.locks"`
	VectorhCpuLoad             MetricSettings `mapstructure:"vectorh.cpuload"`
}


// IsEnabledProvidedByUser returns true if `enabled` option is explicitly set in user settings to any value.


func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		MysqlLocks: MetricSettings{
			Enabled: true,
		},
		VectorhCpuLoad: MetricSettings{
			Enabled: true,
		},

	}
}

// AttributeLocks specifies the a value locks attribute.
type AttributeLocks int

const (
	_ AttributeLocks = iota
	AttributeLocksImmediate
	AttributeLocksWaited
)

// String returns the string representation of the AttributeLocks.
func (av AttributeLocks) String() string {
	switch av {
	case AttributeLocksImmediate:
		return "immediate"
	case AttributeLocksWaited:
		return "waited"
	}
	return ""
}

// MapAttributeLocks is a helper map of string to AttributeLocks attribute value.
var MapAttributeLocks = map[string]AttributeLocks{
	"immediate": AttributeLocksImmediate,
	"waited":    AttributeLocksWaited,
}


//////////////////////////////////////////////////////////////////////////
//  Metric MysqlLocks 
//////////////////////////////////////////////////////////////////////////
type metricMysqlLocks struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills mysql.locks metric with initial data.
func (m *metricMysqlLocks) init() {
	m.data.SetName("mysql.locks")
	m.data.SetDescription("The number of MySQL locks.")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMysqlLocks) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, locksAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutString("kind", locksAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMysqlLocks) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMysqlLocks) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMysqlLocks(settings MetricSettings) metricMysqlLocks {
	m := metricMysqlLocks{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// RecordMysqlLocksDataPoint adds a data point to mysql.locks metric.
func (mb *MetricsBuilder) RecordMysqlLocksDataPoint(ts pcommon.Timestamp, inputVal string, locksAttributeValue AttributeLocks) error {
	val, err := strconv.ParseInt(inputVal, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse int64 for MysqlLocks, value was %s: %w", inputVal, err)
	}
	mb.metricMysqlLocks.recordDataPoint(mb.startTime, ts, val, locksAttributeValue.String())
	return nil
}


//////////////////////////////////////////////////////////////////////////
//  Metric VectorhCpuLoad  
//////////////////////////////////////////////////////////////////////////
type metricVectorhCpuLoad struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills vectorh.cpuLoad metric with initial data.
func (m *metricVectorhCpuLoad ) init() {
	fmt.Println("DEBUG: metricVectorhCpuLoad:   init()")
	m.data.SetName("vectorh.cpuload")
	m.data.SetDescription("Cpu Load")
	m.data.SetUnit("1")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricVectorhCpuLoad) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
    fmt.Println("DEBUG: In recordDataPoint: ")
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutString("kind","cpuLoad" )
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricVectorhCpuLoad) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricVectorhCpuLoad) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricVectorhCpuLoad(settings MetricSettings) metricVectorhCpuLoad{
	m := metricVectorhCpuLoad{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}


// RecordVectorhCpuLoadDataPoint adds a data point to vectorh.cpuload metric.
func (mb *MetricsBuilder) RecordVectorhCpuLoadDataPoint(ts pcommon.Timestamp, inputVal string, errors *scrapererror.ScrapeErrors) error {
	val, err := strconv.ParseFloat(inputVal, 64)
	fmt.Println("DEBUG: In RecordVectorhCpuLoadDataPoint: ",val,inputVal)
	if err != nil {
		fmt.Println("DEBUG: err in RecordVectorhCpuLoadDataPoint")
		return fmt.Errorf("failed to parse int64 for MysqlLocks, value was %s: %w", inputVal, err)
	}
	mb.metricVectorhCpuLoad.recordDataPoint(mb.startTime, ts, val)

	return nil
}


//////////////////////////////////////////////////////////////////////////
//   
//////////////////////////////////////////////////////////////////////////

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                        pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                  int                 // maximum observed number of metrics per resource.
	resourceCapacity                 int                 // maximum observed number of resource attributes.
	metricsBuffer                    pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                        component.BuildInfo // contains version information
	metricMysqlLocks                 metricMysqlLocks
	metricVectorhCpuLoad             metricVectorhCpuLoad
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                        pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                    pmetric.NewMetrics(),
		buildInfo:                        buildInfo,
		metricMysqlLocks:                 newMetricMysqlLocks(settings.MysqlLocks),
		metricVectorhCpuLoad:             newMetricVectorhCpuLoad(settings.VectorhCpuLoad),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithMysqlInstanceEndpoint sets provided value as "mysql.instance.endpoint" attribute for current resource.
func WithMysqlInstanceEndpoint(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutString("mysql.instance.endpoint", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/emareceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricMysqlLocks.emit(ils.Metrics())
	mb.metricVectorhCpuLoad.emit(ils.Metrics())
//	mb.metricMysqlThreads.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}



// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
