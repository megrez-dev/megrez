function dom(id) {
    return document.getElementById(id);
}
function create(tag, attr) {
    var el = document.createElement(tag);
    for (var key in attr) {
        el.setAttribute(key, attr[key]);
    }
    return el;
}
function reply(did, pid, rid, respid) {
    var comment = dom(did),
        response = dom(respid),
        parent = dom('comment-parent'),
        root = dom('comment-root'),
        form = 'form' == response.tagName ? response : response.getElementsByTagName('form')[0],
        textarea = response.getElementsByTagName('textarea')[0];
    if (null == parent) {
        parent = create('input', {
            'type': 'hidden',
            'name': 'parent',
            'id': 'comment-parent'
        });

        form.appendChild(parent);
    }
    parent.setAttribute('value', pid);
    if (null == root) {
        root = create('input', {
            'type': 'hidden',
            'name': 'root',
            'id': 'comment-root'
        });
        form.appendChild(root);
    }
    root.setAttribute('value', rid);
    if (null == dom('comment-form-place-holder')) {
        var holder = create('div', {
            'id': 'comment-form-place-holder'
        });

        response.parentNode.insertBefore(holder, response);
    }
    comment.appendChild(response);
    dom('cancel-comment-reply-link').style.display = '';
    if (null != textarea && 'text' == textarea.name) {
        textarea.focus();
    }
    return false;
}
function replyArticle(aid, cid, rid) {
    var did = 'comment-' + cid;
    var respid = 'respond-post-' + aid;
    reply(did, cid, rid, respid);
}
function replyPage(pid, cid, rid) {
    var did = 'comment-' + cid;
    var respid = 'respond-page-' + pid;
    reply(did, cid, rid, respid);
}
function cancelReply(respid) {
    var response = dom(respid),
        holder = dom('comment-form-place-holder'),
        parent = dom('comment-parent');
    if (null != parent) {
        parent.parentNode.removeChild(parent);
    }
    if (null == holder) {
        return true;
    }
    dom('cancel-comment-reply-link').style.display = 'none';
    holder.parentNode.insertBefore(response, holder);
    return false;
}
function cancelReplyArticle(aid) {
    var respid = 'respond-post-' + aid;
    cancelReply(respid);
}
function cancelReplyPage(pid) {
    var respid = 'respond-page-' + pid;
    cancelReply(respid);
}