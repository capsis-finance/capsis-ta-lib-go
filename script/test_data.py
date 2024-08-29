import argparse
import datetime
import os

import pandas as pd
import pandas_ta as ta
import requests

kline_cols = ['t', 'o', 'h', 'l', 'c', 'vb', 'close_time', 'vq', 'n', 'vbt', 'vqt', 'ignore']
kline_cols_drop = ['close_time', 'ignore', 'vbt', 'vqt', ]


def fetch_data(_symbol, _interval, _limit, _output_csv):
    _data = requests.get('https://api.binance.com/api/v3/klines?symbol=' + _symbol + '&interval=' + _interval + '&limit=' + str(_limit)).json()
    _df = pd.DataFrame(_data, columns=kline_cols).drop(kline_cols_drop, axis=1)
    _df.to_csv(_output_csv, index=False)
    return _df


def get_indicators(_df_kline, _output_csv):
    # copy
    _keep_cols = ['t', 'h', 'l', 'c']
    _df_indicators = _df_kline[_keep_cols].copy()

    # init reset time column
    _df_indicators['t_copy'] = _df_indicators['t'].copy()

    # set time index
    _df_indicators['t'] = pd.to_datetime(_df_indicators['t'], unit='ms') + pd.Timedelta(hours=2)
    _df_indicators = _df_indicators.set_index('t')

    # reset time column
    _df_indicators = _df_indicators.rename(columns={'t_copy': 't'})
    _df_indicators['t'] = _df_indicators['t'].astype('int64')

    # retype
    _skip_cols = ['t']
    for _c in _df_indicators.columns:
        if _c not in _skip_cols:
            _df_indicators[_c] = _df_indicators[_c].astype('float64')

    # bb
    bb = ta.bbands(_df_indicators['c'])
    _df_indicators = _df_indicators.merge(bb[['BBL_5_2.0', 'BBM_5_2.0', 'BBU_5_2.0']], left_index=True, right_index=True, how='outer')

    # ichimoku
    ichimoku_backward, ichimoku_forward = ta.ichimoku(_df_indicators['h'], _df_indicators['l'], _df_indicators['c'])
    # todo: make me var
    ichimoku_forward.index = pd.Series(pd.date_range(ichimoku_backward.index[-1] + pd.Timedelta(minutes=1), freq="min", periods=len(ichimoku_forward.index)))
    ichimoku = pd.concat([ichimoku_backward, ichimoku_forward])
    _df_indicators = _df_indicators.merge(ichimoku, left_index=True, right_index=True, how='outer')

    _df_indicators = _df_indicators.rename(columns={
        'BBL_5_2.0': 'BB_DOWN',
        'BBM_5_2.0': 'BB_MID',
        'BBU_5_2.0': 'BB_UP',
        'ISA_9': 'spanA',
        'ISB_26': 'spanB',
        'ITS_9': 'tenkan',
        'IKS_26': 'kijun',
        'ICS_26': 'chiko',
    })

    # save with index name
    _df_indicators.to_csv(_output_csv, index_label='t')

    return _df_indicators


def main(_symbol, _interval, _kline_limit, _run_output_dir, _export):
    _base_file_name = _run_output_dir + _symbol + '_' + _interval
    print("_base_file_name:", _base_file_name)

    # fetch kline & save
    kline_output = _base_file_name + '_kline.csv'
    df_kline = fetch_data(_symbol, _interval, _kline_limit, kline_output)
    print("kline_output:", kline_output)
    if _export:
        export_kline_path = _run_output_dir + "../test/test_kline.csv"
        df_kline.to_csv(export_kline_path, index=False)
        print("export_kline_path:", export_kline_path)

    # add indicators
    indic_output = _base_file_name + '_indicator.csv'
    df_indic = get_indicators(df_kline, indic_output)
    print("indic_output:", indic_output, "\n")
    print("df_indic:", df_indic)

    if _export:
        export_indic_path = _run_output_dir + "../test/test_indicator.csv"
        df_indic.to_csv(export_indic_path, index_label='t')
        print("export_indic_path:", export_indic_path)


def parse_args():
    # todo
    _parser = argparse.ArgumentParser(description="Download kline and compute indicators (used to populate test_data directory)")
    _parser.add_argument("-s", "--symbol", type=str, help="symbol", required=True)
    _parser.add_argument("-i", "--interval", type=str, help="interval", required=True)
    _parser.add_argument("-l", "--limit", type=int, help="limit", required=False)
    _parser.add_argument("-e", "--export", action='store_true', help="export as test (save to test_data/test)", required=False)
    _args = _parser.parse_args()
    return _args


if __name__ == '__main__':
    # parse args
    args = parse_args()

    if args.limit is None:
        args.limit = 500

    # init output name
    _base_output_name = args.symbol + '_' + args.interval + "_limit" + str(args.limit)

    # init output dir
    run_time = datetime.datetime.now().strftime("%Y_%m_%d_%H:%M:%S")
    run_output_dir = "./test_data/run_" + run_time + "-" + _base_output_name + "/"
    print("run_output_dir:", run_output_dir)
    os.mkdir(run_output_dir)

    main(args.symbol, args.interval, args.limit, run_output_dir, args.export)
