import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import relativeTime from 'dayjs/plugin/relativeTime';

dayjs.locale('zh-cn');
dayjs.extend(relativeTime);


function datetimeFormat(datetime = new Date(), pattern = 'YYYY-MM-DD HH:mm') {
    return dayjs(datetime).format(pattern)
}

function timeAgo(datetime) {
    if (dayjs().diff(dayjs(datetime), 'd') < 5) return dayjs(datetime).fromNow();
    return dayjs(datetime).format('YYYY-MM-DD HH:mm');
}

export { datetimeFormat, timeAgo }
