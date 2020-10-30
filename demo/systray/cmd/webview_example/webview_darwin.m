#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>
#include "webview.h"

NSWindowController *windowController = nil;
NSWindow *Window = nil;
WKWebView *webView = nil;

void configureAppWindow(char* title, int width, int height)
{
  if (windowController != nil) {
    // already configured, ignore
    return;
  }

  NSApplication *app = [NSApplication sharedApplication];
  [app setActivationPolicy:NSApplicationActivationPolicyRegular];
  [app activateIgnoringOtherApps:YES];

  NSRect frame = NSMakeRect(0, 0, width, height);
  int mask = NSWindowStyleMaskTitled | NSWindowStyleMaskResizable | NSWindowStyleMaskClosable;
  Window = [[NSWindow alloc] initWithContentRect:frame
                              styleMask:mask
                              backing:NSBackingStoreBuffered
                              defer:NO];
  [Window setTitle:[[NSString alloc] initWithUTF8String:title]];
  [Window center];

  NSView *contentView = [Window contentView];
  webView = [[WKWebView alloc] initWithFrame:[contentView bounds]];
  [webView setTranslatesAutoresizingMaskIntoConstraints:NO];
  [contentView addSubview:webView];
  [contentView addConstraint:
    [NSLayoutConstraint constraintWithItem:webView
        attribute:NSLayoutAttributeWidth
        relatedBy:NSLayoutRelationEqual
        toItem:contentView
        attribute:NSLayoutAttributeWidth
        multiplier:1
        constant:0]];
  [contentView addConstraint:
    [NSLayoutConstraint constraintWithItem:webView
        attribute:NSLayoutAttributeHeight
        relatedBy:NSLayoutRelationEqual
        toItem:contentView
        attribute:NSLayoutAttributeHeight
        multiplier:1
        constant:0]];

  // Window controller:
  windowController = [[NSWindowController alloc] initWithWindow:Window];
  
  free(title);
  [NSApp run];
}

void doShowAppWindow(char* url)
{
  if (windowController == nil) {
    // no app Window to open
    return;
  }

  id nsURL = [NSURL URLWithString:[[NSString alloc] initWithUTF8String:url]];
  id req = [[NSURLRequest alloc] initWithURL: nsURL
                                 cachePolicy: NSURLRequestUseProtocolCachePolicy
                                 timeoutInterval: 5];
  [webView loadRequest:req];
  [windowController showWindow:Window];
  free(url);
}

void showAppWindow(char* url)
{
  dispatch_async(dispatch_get_main_queue(), ^{
    doShowAppWindow(url);
  });
}
